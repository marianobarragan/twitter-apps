package domain

import (
	"time"
)

type Timeline struct {
	Tweets []TimelineTweet `json:"timelineTweets"`
}

type TimelineTweet struct {
	ID        int       `json:"id"`
	TweetID   int       `json:"tweet_id"`
	Text      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	AuthorID  int       `json:"author_id"`
	Follower  int       `json:"follower_id"`
}

type Tweet struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
