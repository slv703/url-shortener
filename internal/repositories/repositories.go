package repositories

import "log/slog"

type Repositories struct {
	URL *URLRepository
}

func NewRepositories(logger *slog.Logger) *Repositories {
	return &Repositories{
		URL: NewURLRepository(logger),
	}
}
