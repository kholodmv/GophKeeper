package handlers

import (
	"github.com/kholodmv/GophKeeper/cmd/server/config"
	"github.com/kholodmv/GophKeeper/internal/services/auth"
	"github.com/kholodmv/GophKeeper/internal/services/keeper"
	"github.com/kholodmv/GophKeeper/internal/storage"
	"golang.org/x/exp/slog"
	"os"
)

type Handler struct {
	Config        *config.Config
	Storage       *storage.Storage
	AuthService   *auth.Auth
	Log           *slog.Logger
	KeeperService *keeper.Keeper
}

func NewHandler(config *config.Config) *Handler {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	store := storage.NewStorage(config.DatabaseURI)

	return &Handler{
		Config:        config,
		Storage:       store,
		Log:           log,
		AuthService:   auth.NewAuth(log, store, config.TokenTTL),
		KeeperService: keeper.NewKeeper(log, store),
	}
}
