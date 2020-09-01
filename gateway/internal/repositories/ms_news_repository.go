package repositories

import (
	"context"

	"gateway/internal/models"
	"gateway/pkg/pb"
	"github.com/golang/protobuf/ptypes/empty"
)

type msNewsRepository struct {
	newsGrpc pb.NewsServiceClient
}

func NewMSNewsRepository(newsGrpc pb.NewsServiceClient) *msNewsRepository {
	return &msNewsRepository{
		newsGrpc: newsGrpc,
	}
}

func (repo *msNewsRepository) Get(ctx context.Context) ([]models.News, error) {
	newsListResponse, err := repo.newsGrpc.GetNews(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}

	news := make([]models.News, 0)
	for _, item := range newsListResponse.News {
		news = append(news, models.News{
			ID:     uint(item.Id),
			Title:  item.Title,
			Body:   item.Body,
			UserID: uint(item.UserId),
		})
	}

	return news, nil
}
