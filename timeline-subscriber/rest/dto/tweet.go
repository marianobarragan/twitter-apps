package dto

import "time"

type Tweet struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	AuthorID  int       `json:"author_id"`
	CreatedAt time.Time `json:"created_at"`
}
