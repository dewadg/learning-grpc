package services

import (
	"context"
	"sync"

	"gateway/internal/models"
	"gateway/internal/repositories"
)

type newsService struct {
	userRepo repositories.UserRepository
	newsRepo repositories.NewsRepository
}

func NewNewsService(userRepo repositories.UserRepository, newsRepo repositories.NewsRepository) *newsService {
	return &newsService{
		userRepo: userRepo,
		newsRepo: newsRepo,
	}
}

func (svc *newsService) Get(ctx context.Context) ([]models.News, error) {
	news, err := svc.newsRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	userMap, err := svc.buildUserMap(ctx, news)
	if err != nil {
		return nil, err
	}

	for i := range news {
		if user, ok := userMap[news[i].UserID]; ok {
			news[i].User = &user
		}
	}

	return news, err
}

func (svc *newsService) extractUserIDsFromNews(news []models.News) []uint {
	hashMap := make(map[uint]bool)
	result := make([]uint, 0)

	for _, item := range news {
		if _, ok := hashMap[item.UserID]; ok {
			continue
		}

		hashMap[item.UserID] = true
		result = append(result, item.UserID)
	}

	return result
}

func (svc *newsService) buildUserMap(ctx context.Context, news []models.News) (map[uint]models.User, error) {
	userIDs := svc.extractUserIDsFromNews(news)
	userChan := make(chan models.User)
	errChan := make(chan error)
	doneChan := make(chan bool)
	wg := sync.WaitGroup{}

	for _, userID := range userIDs {
		wg.Add(1)

		go func(wg *sync.WaitGroup, userID uint) {
			defer wg.Done()

			user, err := svc.userRepo.Find(ctx, userID)
			if err != nil {
				errChan <- err
				return
			}

			userChan <- user
		}(&wg, userID)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()

		doneChan <- true
	}(&wg)

	userMap := make(map[uint]models.User)
	for {
		select {
		case user := <-userChan:
			userMap[user.ID] = user
		case err := <-errChan:
			return nil, err
		case <-doneChan:
			return userMap, nil
		}
	}
}
