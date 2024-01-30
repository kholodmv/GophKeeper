package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

var Log *slog.Logger

func init() {
	// Initializing the logger at program startup
	Log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
}
