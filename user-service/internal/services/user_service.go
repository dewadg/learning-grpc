package services

import (
	"context"

	"user-service/internal/models"
	"user-service/internal/repositories"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{
		userRepo: userRepo,
	}
}

func (svc *userService) Find(ctx context.Context, id uint) (models.User, error) {
	return svc.userRepo.Find(ctx, id)
}
