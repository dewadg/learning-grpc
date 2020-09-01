package grpc

import (
	"context"

	"user-service/internal/services"
	"user-service/pkg/pb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userSvc services.UserService
}

func NewUserHandler(userSvc services.UserService) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (h *UserHandler) FindUser(ctx context.Context, input *pb.FindUserInput) (*pb.User, error) {
	user, err := h.userSvc.Find(ctx, uint(input.Id))
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:       uint32(user.ID),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
