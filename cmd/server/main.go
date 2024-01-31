package main

import (
	"context"
	"github.com/kholodmv/GophKeeper/cmd/server/config"
	"github.com/kholodmv/GophKeeper/internal/handlers"
	"github.com/kholodmv/GophKeeper/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.UseServerStartParams()
	rootCtx := context.Background()
	handler := handlers.NewHandler(cfg, rootCtx)
	server := &http.Server{
		Addr:    cfg.RunAddress,
		Handler: router.Router(handler),
	}

	connectionsClosed := make(chan struct{})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		<-stop
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}

		close(connectionsClosed)
	}()

	run(server)
}

func run(server *http.Server) {
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}
