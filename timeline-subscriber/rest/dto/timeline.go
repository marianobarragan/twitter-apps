package dto

import "time"

type Timeline struct {
	AuthorID int `json:"author_id"`
	Tweet    struct {
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"tweet"`
}
