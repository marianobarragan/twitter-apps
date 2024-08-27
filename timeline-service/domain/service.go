package domain

import "time"

type Service interface {
	GetUserTimeline(id int, from time.Time, to time.Time) (Timeline, error)
	AddTweetToUserTimeline(userID int, tweet Tweet) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return service{
		repository: r,
	}
}

func (ps service) GetUserTimeline(id int, from time.Time, to time.Time) (Timeline, error) {
	return ps.repository.GetUserTimeline(id, from, to)
}

func (ps service) AddTweetToUserTimeline(userID int, tweet Tweet) error {
	return ps.repository.AddTweetToUserTimeline(userID, tweet)
}
