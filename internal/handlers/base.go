package handlers

import (
	"log/slog"

	"github.com/slv703/url-shortener/internal/repositories"
)

type BaseHandler struct {
	logger *slog.Logger
	repositories *repositories.Repositories
}

func NewBaseHandler(logger *slog.Logger, name string, repositories *repositories.Repositories) *BaseHandler {
	return &BaseHandler{logger: logger.With("Handler", name), repositories: repositories}
}
