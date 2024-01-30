package handlers

import (
	"context"
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
	Ctx           context.Context
}

func NewHandler(config *config.Config, ctx context.Context) *Handler {
	store := storage.NewStorage(config.DatabaseURI)

	return &Handler{
		Config:        config,
		Storage:       store,
		AuthService:   auth.NewAuth(store, config.TokenTTL),
		KeeperService: keeper.NewKeeper(store),
		Ctx:           ctx,
	}
}
