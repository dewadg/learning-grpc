package grpc

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"news-service/internal/services"
	"news-service/pkg/pb"
)

type NewsHandler struct {
	pb.UnimplementedNewsServiceServer
	newsSvc services.NewsService
}

func NewNewsHandler(newsSvc services.NewsService) *NewsHandler {
	return &NewsHandler{
		newsSvc: newsSvc,
	}
}

func (h *NewsHandler) GetNews(ctx context.Context, input *empty.Empty) (*pb.NewsListResponse, error) {
	news, err := h.newsSvc.Get(ctx)
	if err != nil {
		return nil, err
	}

	response := &pb.NewsListResponse{
		News: make([]*pb.News, 0),
	}

	for _, n := range news {
		response.News = append(response.News, &pb.News{
			Id:     uint32(n.ID),
			Title:  n.Title,
			Body:   n.Body,
			UserId: uint32(n.UserID),
		})
	}

	return response, nil
}
