package domain

import (
	"time"
	"timeline-subscriber/rest"
	"timeline-subscriber/rest/dto"
)

type Service interface {
	IndexTweetTimeline(tweetID int) error
}

type service struct {
	timelineClient rest.TimelineClient
	tweetsClient   rest.TweetsClient
	usersClient    rest.UsersClient
}

func NewService(timelineClient rest.TimelineClient, tweetsClient rest.TweetsClient, usersClient rest.UsersClient) Service {
	return service{
		timelineClient: timelineClient,
		tweetsClient:   tweetsClient,
		usersClient:    usersClient,
	}
}

func (ps service) IndexTweetTimeline(tweetID int) error {
	tweet, err := ps.tweetsClient.GetTweet(tweetID)
	if err != nil {
		return err
	}

	user, err := ps.usersClient.GetUser(tweet.AuthorID)
	if err != nil {
		return err
	}

	for _, follower := range user.Followers {
		err := ps.updateFeed(follower, tweet)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ps service) updateFeed(follower dto.Follower, tweet dto.Tweet) error {
	timeline := dto.Timeline{
		AuthorID: tweet.AuthorID,
		Tweet: struct {
			Text      string    `json:"text"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Text:      tweet.Text,
			CreatedAt: tweet.CreatedAt,
		},
	}
	err := ps.timelineClient.AddTimeline(follower.UserID, timeline)
	return err
}
