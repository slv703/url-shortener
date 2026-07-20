package repositories

import (
	"context"
	"log/slog"

	"github.com/slv703/url-shortener/internal/models"
)

type URLRepository struct {
	BaseRepository
}

func NewURLRepository(logger *slog.Logger) *URLRepository {
	return &URLRepository{BaseRepository: *NewBaseRepository(logger)}
}

func (r *URLRepository) FindByKey(ctx context.Context, key string) (*models.URL, error) {
	data := map[string]models.URL{
		"3UsWTj": {
			Key:         "3UsWTj",
			OriginalURL: "https://www.ozon.ru/product/svetocopy-a4-500-l-bumaga-dlya-printera-7969279/",
		},
	}

	url, ok := data[key]
	if !ok {
		return nil, ErrRecordNotFound
	}
	return &url, nil
}
