package events

type NewTweetEvent struct {
	TweetID  int `json:"tweet_id"`
	AuthorID int `json:"author_id"`
}
