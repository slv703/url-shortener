package internal

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/slv703/url-shortener/internal/config"
	"github.com/slv703/url-shortener/internal/handlers"
	"github.com/slv703/url-shortener/internal/repositories"
)

type Application struct {
	config       *config.Config
	logger       *slog.Logger
	repositories *repositories.Repositories
	server       *echo.Echo
}

func NewApplication(cfg *config.Config, logger *slog.Logger) *Application {
	app := Application{
		config: cfg,
		logger: logger,
		repositories: repositories.NewRepositories(logger),
		server: setupServer(cfg.Server, logger),
	}
	app.setupRoutes()
	return &app
}

func (a *Application) Start(ctx context.Context) {
	go func() {
		err := a.server.Start(net.JoinHostPort(a.config.Server.Host, a.config.Server.Port))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.logger.Error(fmt.Sprintf("start server error: %s", err.Error()))
		}
	}()
}

func (a *Application) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	a.logger.Info("Stopping web server...")
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error(fmt.Sprintf("web server shutdown error: %s", err.Error()))
	}
	a.logger.Info("Web server stopped")
}

func (a *Application) setupRoutes() {
	urlHandler := handlers.NewURLHandler(a.logger, a.repositories)

	a.server.GET("/hello", a.helloHandler)
	a.server.GET("/:key", urlHandler.Redirect)
}

func (a *Application) helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
