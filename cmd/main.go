package main

import (
	"log/slog"
	"os"
)

func initLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return log
}

func main() {
	l := initLogger()

	l.Info("starting...")
}
