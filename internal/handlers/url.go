package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/slv703/url-shortener/internal/repositories"
)

type URLHandler struct {
	*BaseHandler
}

func NewURLHandler(logger *slog.Logger, repositories *repositories.Repositories) *URLHandler {
	return &URLHandler{BaseHandler: NewBaseHandler(logger, "URLHandler", repositories)}
}

func (h *URLHandler) Redirect(c echo.Context) error {
	logger := h.logger.With("Op", "Redirect")

	key := c.Param("key")
	logger.Info("Request", slog.String("Key", key))

	url, err := h.repositories.URL.FindByKey(c.Request().Context(), key)
	if err != nil && errors.Is(err, repositories.ErrRecordNotFound) {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.Redirect(http.StatusFound, url.OriginalURL)
}
