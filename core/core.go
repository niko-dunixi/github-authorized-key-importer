package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Key struct {
	ID  int    `json:"id"`
	Key string `json:"key"`
}

func GetKeys(username string) ([]Key, error) {
	requestURL := url.URL{
		Scheme: "https",
		Host:   "api.github.com",
		Path:   "users/" + username + "/keys",
	}
	request, err := http.NewRequest("GET", requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't form request: %w", err)
	}

	client := HttpClient()
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("couldn't retrieve keys: %w", err)
	}
	keys := make([]Key, 0, 10)
	if err := json.NewDecoder(response.Body).Decode(&keys); err != nil {
		return nil, fmt.Errorf("could not parse keys: %w", err)
	}
	return keys, nil
}

func HttpClient() http.Client {
	return http.Client{}
}
