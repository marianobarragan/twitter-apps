package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"timeline-subscriber/rest/dto"
)

type TimelineClient interface {
	AddTimeline(userID int, timeline dto.Timeline) error
}

func NewTimelineClient() TimelineClient {
	return RealTimelineClient{}
}

type RealTimelineClient struct{}

func (TimelineClient RealTimelineClient) AddTimeline(userID int, timeline dto.Timeline) error {
	requestURL := fmt.Sprintf("http://localhost:%d/users/%d/timeline", serverPort, userID)
	requestBody, err := json.Marshal(timeline)
	if err != nil {
		err = fmt.Errorf("failed to marshal request data: %v", err)
		return err
	}
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("error making http request to %s: %s\n", requestURL, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("received non-201 response code: %d", resp.StatusCode)
		return err
	}
	return nil
}
