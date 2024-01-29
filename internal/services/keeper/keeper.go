package keeper

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/models"
	"golang.org/x/exp/slog"
)

type Keeper struct {
	log      *slog.Logger
	provider SecretProvider
}

// SecretProvider - type, specifying where secrets are create and read.
type SecretProvider interface {
	CreateSecret(secret *models.Secret) error
	ReadSecret(title string, uid uint) (*models.Secret, error)
}

// NewKeeper - keeper constructor
func NewKeeper(
	log *slog.Logger,
	provider SecretProvider,
) *Keeper {
	return &Keeper{
		provider: provider,
		log:      log,
	}
}

// CreateSecret - create new secret
func (k *Keeper) CreateSecret(secret *models.Secret) error {
	k.log.Info("create secret")
	err := k.provider.CreateSecret(secret)
	if err != nil {
		k.log.Error("Keeper.Create: ", err)
		return fmt.Errorf("%s: %w", "Keeper.Create: ", err)
	}
	return nil
}

func (k *Keeper) ReadSecret(title string, uid uint) (*models.Secret, error) {
	k.log.Info("read secret")
	secret, err := k.provider.ReadSecret(title, uid)
	if err != nil {
		k.log.Error("Keeper.Read: ", err)

		return nil, fmt.Errorf("%s: %w", "Keeper.Create: ", err)
	}

	return secret, nil
}
