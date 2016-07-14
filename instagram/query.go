package instagram

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	gu = `/query/?q=ig_user(%d){username}`
)

// QueryService - сервис работы с запросами query
type QueryService struct {
	client *Client
}

// GetUsernameByID - получение профиля пользователя по ID
func (s *QueryService) GetUsernameByID(id uint) (username string, err error) {

	data, err := s.get(gu, id)

	if _, ok := data["username"]; !ok {
		return username, errors.New("User not found")
	}

	return data["username"].(string), err
}

func (s *QueryService) get(query string, params ...interface{}) (data map[string]interface{}, err error) {

	u := fmt.Sprintf(query, params...)

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &data)

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New("Error in query")
	}

	if _, ok := data["ok"]; !ok {
		return data, err
	}

	return data, err
}
