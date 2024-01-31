package keeper

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/logger"
	"github.com/kholodmv/GophKeeper/internal/models"
)

type Keeper struct {
	provider SecretProvider
}

// SecretProvider - type, specifying where secrets are create and read.
type SecretProvider interface {
	CreateSecret(secret *models.Secret) error
	ReadSecret(title string, uid uint) (*models.Secret, error)
}

// NewKeeper - keeper constructor
func NewKeeper(
	provider SecretProvider,
) *Keeper {
	return &Keeper{
		provider: provider,
	}
}

// CreateSecret - create new secret
func (k *Keeper) CreateSecret(secret *models.Secret) error {
	logger.Log.Info("create secret")
	err := k.provider.CreateSecret(secret)
	if err != nil {
		logger.Log.Error("Keeper.Create: ", err)
		return fmt.Errorf("%s: %w", "Keeper.Create: ", err)
	}
	return nil
}

func (k *Keeper) ReadSecret(title string, uid uint) (*models.Secret, error) {
	logger.Log.Info("read secret")
	secret, err := k.provider.ReadSecret(title, uid)
	if err != nil {
		logger.Log.Error("Keeper.Read: ", err)

		return nil, fmt.Errorf("%s: %w", "Keeper.Create: ", err)
	}

	return secret, nil
}
