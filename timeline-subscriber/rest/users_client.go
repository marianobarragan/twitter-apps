package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"timeline-subscriber/rest/dto"
)

const serverPort = 8083

type UsersClient interface {
	GetUser(userID int) (dto.User, error)
}

func NewUsersClient() UsersClient {
	return RealUsersClient{}
}

type RealUsersClient struct{}

func (usersClient RealUsersClient) GetUser(userID int) (user dto.User, err error) {
	requestURL := fmt.Sprintf("http://localhost:%d/users/%d", serverPort, userID)
	resp, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request to %s: %s\n", requestURL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("failed to read response body")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		err = errors.New("failed to parse JSON response")
		return
	}

	return
}
