package services

import (
	"context"

	"news-service/internal/models"
)

type NewsService interface {
	Get(ctx context.Context) ([]models.News, error)
}
