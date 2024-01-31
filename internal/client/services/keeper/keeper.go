package keeper

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/kholodmv/GophKeeper/internal/client/models"
	"github.com/kholodmv/GophKeeper/internal/client/storage"
	"net/http"
	"strings"
)

// KeeperService is an interface that provides methods for authentication and registration.
type KeeperService interface {
	Create(secret *models.Secret) error
	Get(title string) (*models.Secret, error)

	// GetClient returns the service's client.
	GetClient() *resty.Client
}

// keeperService is a concrete implementation of KeeperService.
type keeperService struct {
	client *resty.Client
}

// newConfiguredClient returns a client configured for https (if required).
func newConfiguredClient(baseURL string) *resty.Client {
	client := resty.New().SetBaseURL(baseURL)
	if strings.Contains(baseURL, "https") {
		client = client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	return client
}

// NewKeeperService creates a new instance of authService with the given baseURL and returns it as an AuthService.
func NewKeeperService(baseURL string) KeeperService {
	client := newConfiguredClient(baseURL)
	return &keeperService{client: client}
}

func (k *keeperService) Create(secret *models.Secret) error {
	r := &models.CreateResponse{}

	tokenStorage := storage.New()
	token := tokenStorage.GetToken()

	resp, err := k.client.R().
		SetResult(r).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(secret).
		Post("/api/secret/create")
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return errors.New(r.Message)
	}
	return nil
}

// Get authenticates a user with the given username and password and returns an authentication token if successful.
func (k *keeperService) Get(title string) (*models.Secret, error) {
	r := &models.Secret{}

	tokenStorage := storage.New()
	token := tokenStorage.GetToken()

	resp, err := k.client.R().
		SetResult(r).
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", token).
		SetBody(fmt.Sprintf(`{"title":"%s"}`, title)).
		Post("/api/secret/read")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, err
	}

	return r, nil
}

// GetClient returns the service's client.
func (k *keeperService) GetClient() *resty.Client {
	return k.client
}
