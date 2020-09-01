package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"user-service/internal/models"
)

type jsonPlaceholderUserRepository struct {
	client *http.Client
}

func NewJSONPlaceholderUserRepository() *jsonPlaceholderUserRepository {
	return &jsonPlaceholderUserRepository{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (repo *jsonPlaceholderUserRepository) Find(ctx context.Context, id uint) (models.User, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.User{}, err
	}

	response, err := repo.client.Do(request)
	if err != nil {
		return models.User{}, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return models.User{}, fmt.Errorf("%d: request_failed_with_status", response.StatusCode)
	}

	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = json.Unmarshal(responseByte, &user)

	return user, err
}
