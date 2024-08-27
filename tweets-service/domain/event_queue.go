package domain

type EventProducer interface {
	PublishNewTweetEvent(userID int, tweet Tweet) error
	Close() error
}
