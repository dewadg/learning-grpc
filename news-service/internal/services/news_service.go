package services

import (
	"context"

	"news-service/internal/models"
	"news-service/internal/repositories"
)

type newsService struct {
	repo repositories.NewsRepository
}

func NewNewsService(repo repositories.NewsRepository) *newsService {
	return &newsService{
		repo: repo,
	}
}

func (svc *newsService) Get(ctx context.Context) ([]models.News, error) {
	return svc.repo.Get(ctx)
}
