package services

import (
	"context"

	"gateway/internal/models"
)

type NewsService interface {
	Get(ctx context.Context) ([]models.News, error)
}
