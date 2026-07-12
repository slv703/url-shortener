package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/slv703/url-shortener/internal"
	"github.com/slv703/url-shortener/internal/config"
)

func main() {
	// I. Init context
	ctx := context.Background()

	// II. Load config
	config, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("load config error: %s", err.Error())
	}

	// III. Init logger
	logger := initLogger(config.Env)
	logger.Info("Init application", slog.Any("config", config))

	// IV. Init application
	app := internal.NewApplication(config, logger)

	// V. Start application
	app.Start(ctx)

	// VI. Application grecefully shut down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Info(fmt.Sprintf("%s signal was received", <-quit))
	app.Stop()
	logger.Info("Application has shutting down gracefully")
}

func initLogger(env string) *slog.Logger {
	switch env {
	case "local":
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	default:
		return slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
}
