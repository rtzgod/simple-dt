package main

import (
	"github.com/rtzgod/logger"
	"github.com/rtzgod/simple-dt/internal/app/httpserver"
	"github.com/rtzgod/simple-dt/internal/config"
	"github.com/rtzgod/simple-dt/internal/handler"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	handlers := handler.NewHandler()

	server := httpserver.NewServer(log, cfg.HTTP.Port, cfg.HTTP.ReadTimeout, cfg.HTTP.WriteTimeout, handlers.InitRoutes())

	go server.MustRun()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sign := <-sigChan

	log.Info("Stopping server", slog.String("signal:", sign.String()))

	if err := server.Stop(); err != nil {
		log.Error("Failed to stop server", logger.Err(err))
	}

	log.Info("Server stopped")
}
