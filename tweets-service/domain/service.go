package domain

import "log"

type Service interface {
	GetTweet(id int) (Tweet, bool, error)
	SaveTweet(userID int, tweet Tweet) (Tweet, error)
}

type service struct {
	repository Repository
	eventQueue EventProducer
}

func NewService(r Repository, e EventProducer) Service {
	return service{
		repository: r,
		eventQueue: e,
	}
}

func (ps service) GetTweet(id int) (Tweet, bool, error) {
	return ps.repository.GetTweet(id)
}

func (ps service) SaveTweet(userID int, tweet Tweet) (Tweet, error) {
	tweet.AuthorID = userID
	tweet, tx, err := ps.repository.PostTweet(tweet)

	// We have a dual-write problem here: we want to write to the DB and publish the event in an
	//atomic fashion.
	// For this exercise I've fixed this issue by ensuring that the event is emitted before
	//committing the transaction
	err = ps.eventQueue.PublishNewTweetEvent(userID, tweet)
	if err != nil {
		tx.Abort()
	}
	tx.Commit()
	log.Println("Published a new event")
	return tweet, err
}
