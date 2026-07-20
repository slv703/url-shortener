package repositories

import (
	"errors"
	"log/slog"
)

var ErrRecordNotFound = errors.New("record not found")

type BaseRepository struct {
	logger *slog.Logger
}

func NewBaseRepository(logger *slog.Logger) *BaseRepository {
	return &BaseRepository{logger: logger}
}
