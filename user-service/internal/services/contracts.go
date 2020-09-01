package services

import (
	"context"

	"user-service/internal/models"
)

type UserService interface {
	Find(ctx context.Context, id uint) (models.User, error)
}
