package main

import (
	"log/slog"
	"os"

	"github.com/aidosgal/neuron/app"
	"github.com/aidosgal/neuron/config"
)

func main() {
	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		),
	)

	log.Info("starting application")

	application := app.New(cfg, log)

	if err := application.Run(); err != nil {
		panic(err)
	}
}
