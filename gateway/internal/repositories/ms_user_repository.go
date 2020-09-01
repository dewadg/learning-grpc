package repositories

import (
	"context"

	"gateway/internal/models"
	"gateway/pkg/pb"
)

type msUserRepository struct {
	userGrpc pb.UserServiceClient
}

func NewMSUserRepository(userGrpc pb.UserServiceClient) *msUserRepository {
	return &msUserRepository{
		userGrpc: userGrpc,
	}
}

func (repo *msUserRepository) Find(ctx context.Context, id uint) (models.User, error) {
	userResponse, err := repo.userGrpc.FindUser(ctx, &pb.FindUserInput{Id: uint32(id)})
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:       uint(userResponse.Id),
		Username: userResponse.Username,
		Email:    userResponse.Email,
	}, nil
}
