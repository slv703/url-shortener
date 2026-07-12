package internal

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/slv703/url-shortener/internal/config"
)

func setupServer(cfg *config.Server, logger *slog.Logger) *echo.Echo {
	e := echo.New()
	e.Use(timeoutMiddleWare(cfg.Timeout))
	e.Use(middleware.RequestID())
	return e
}

func timeoutMiddleWare(timeout time.Duration) echo.MiddlewareFunc {
	return middleware.ContextTimeoutWithConfig(middleware.ContextTimeoutConfig{Timeout: timeout})
}

func requestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
