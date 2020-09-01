package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"news-service/internal/models"
)

type jsonPlaceholderNewsRepository struct {
	client *http.Client
}

func NewJSONPlaceholderNewsRepository() *jsonPlaceholderNewsRepository {
	return &jsonPlaceholderNewsRepository{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (repo *jsonPlaceholderNewsRepository) Get(ctx context.Context) ([]models.News, error) {
	url := "https://jsonplaceholder.typicode.com/posts"
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := repo.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: request_failed_with_status", response.StatusCode)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var news []models.News
	err = json.Unmarshal(responseByte, &news)

	return news, err
}
