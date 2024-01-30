package auth

import (
	"errors"
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/logger"
	"github.com/kholodmv/GophKeeper/internal/models"
	"github.com/kholodmv/GophKeeper/internal/utils/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// Auth - registration and authentication service
type Auth struct {
	provider UserProvider
	tokenTTL time.Duration
}

// UserProvider - user provider
type UserProvider interface {
	CreateUser(user *models.User) error
	ReadUser(login string) (*models.User, error)
}

var ErrInvalidCredentials = errors.New("invalid credentials")

// NewAuth - auth constructor
func NewAuth(
	provider UserProvider,
	tokenTTL time.Duration,
) *Auth {
	return &Auth{
		provider: provider,
		tokenTTL: tokenTTL, // Lifetime of returned tokens
	}
}

// CreateUser - create new user
func (a *Auth) CreateUser(auth *models.Auth) error {
	logger.Log.Info("registering user")

	passHash, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("failed to generate password hash", err)
		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser.CreatePassHash", err)
	}

	user := &models.User{
		Login:    auth.Login,
		Password: passHash,
	}

	err = a.provider.CreateUser(user)

	if err != nil {
		logger.Log.Error("Auth.RegisterUser: ", err)
		return fmt.Errorf("%s: %w", "Auth.RegisterNewUser", err)
	}

	return nil
}

// Login - user authentication
func (a *Auth) Login(auth *models.Auth) (string, error) {
	logger.Log.Info("login user")
	// Getting the user from the database
	user, err := a.provider.ReadUser(auth.Login)
	if err != nil {
		logger.Log.Error("Auth.Login: ", err)
		return "", fmt.Errorf("%s: %w", "Auth.Login: ", err)
	}

	// Checking the correctness of the received password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(auth.Password)); err != nil {
		logger.Log.Info("invalid credentials", err)
		return "", fmt.Errorf("%s: %w", "Auth.Login: ", ErrInvalidCredentials)
	}

	logger.Log.Info("user logged in successfully")

	// Create an authorization token
	token, err := jwt.NewToken(user, a.tokenTTL)
	if err != nil {
		logger.Log.Error("failed to generate token", err)

		return "", fmt.Errorf("%s: %w", "Auth.Login: ", err)
	}

	return token, nil
}
