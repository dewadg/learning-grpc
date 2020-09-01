package repositories

import (
	"context"

	"user-service/internal/models"
)

type UserRepository interface {
	Find(ctx context.Context, id uint) (models.User, error)
}
