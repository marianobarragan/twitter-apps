package dto

import "time"

type User struct {
	ID        int         `json:"id"`
	Name      string      `json:"name"`
	CreatedAt time.Time   `json:"created_at"`
	Following []Following `json:"followings"`
	Followers []Follower  `json:"followers"`
}

type Following struct {
	UserID int `json:"user_id"`
}
type Follower struct {
	UserID int `json:"user_id"`
}
