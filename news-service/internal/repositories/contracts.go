package repositories

import (
	"context"

	"news-service/internal/models"
)

type NewsRepository interface {
	Get(ctx context.Context) ([]models.News, error)
}
