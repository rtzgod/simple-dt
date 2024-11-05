package main

import (
	"github.com/rtzgod/logger"
	"github.com/rtzgod/simple-dt/internal/app/httpclient"
	"github.com/rtzgod/simple-dt/internal/config"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	client := httpclient.NewClient(log, cfg.Client.Url)

	if err := client.SetData("test"); err != nil {
		log.Error("Failed to set data", logger.Err(err))
	}
}
