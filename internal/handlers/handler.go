package handlers

import (
	"github.com/kholodmv/GophKeeper/cmd/server/config"
	"github.com/kholodmv/GophKeeper/internal/services/auth"
	"github.com/kholodmv/GophKeeper/internal/services/keeper"
	"github.com/kholodmv/GophKeeper/internal/storage"
)

type Handler struct {
	Config        *config.Config
	Storage       *storage.Storage
	AuthService   *auth.Auth
	KeeperService *keeper.Keeper
}

func NewHandler(config *config.Config) *Handler {
	store := storage.NewStorage(config.DatabaseURI)

	return &Handler{
		Config:        config,
		Storage:       store,
		AuthService:   auth.NewAuth(store, config.TokenTTL),
		KeeperService: keeper.NewKeeper(store),
	}
}
