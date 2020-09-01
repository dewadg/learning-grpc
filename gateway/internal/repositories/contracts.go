package repositories

import (
	"context"

	"gateway/internal/models"
)

type NewsRepository interface {
	Get(ctx context.Context) ([]models.News, error)
}

type UserRepository interface {
	Find(ctx context.Context, id uint) (models.User, error)
}
