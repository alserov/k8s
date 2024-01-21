package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"
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

	printText()
}

func printText() {
	for {
		time.Sleep(time.Second)
		fmt.Println("Printing...")
	}
}
