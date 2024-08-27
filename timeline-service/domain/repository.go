package domain

import (
	"time"
)

type Repository interface {
	GetUserTimeline(id int, from time.Time, to time.Time) (Timeline, error)
	AddTweetToUserTimeline(userID int, tweet Tweet) error
}
