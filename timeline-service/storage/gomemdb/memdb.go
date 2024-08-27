package gomemdb

import (
	"github.com/hashicorp/go-memdb"
	"time"
	"timeline-service/domain"
)

type RealRepository struct {
	db *memdb.MemDB
}

var _ domain.Repository = RealRepository{}

func NewRepository() (RealRepository, error) {
	db, err := memdb.NewMemDB(CreateSchema())
	if err != nil {
		return RealRepository{}, err
	}

	return RealRepository{db}, nil
}

func CreateSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"timelineTweets": &memdb.TableSchema{
				Name: "timelineTweets",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
					"follower": &memdb.IndexSchema{
						Name:    "follower",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "Follower"},
					},
				},
			},
		},
	}
}

func (r RealRepository) GetUserTimeline(id int, from time.Time, to time.Time) (domain.Timeline, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	// Here we search for the tweets of the user followings
	// We filter only those that were created during the given time window
	it, err := txn.Get("timelineTweets", "id")
	if err != nil {
		return domain.Timeline{}, err
	}

	timeline := domain.Timeline{
		Tweets: make([]domain.TimelineTweet, 0),
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		t, ok := obj.(*domain.TimelineTweet)
		if ok && t.CreatedAt.After(from) && t.CreatedAt.Before(to) && t.Follower == id {
			timeline.Tweets = append(timeline.Tweets, *t)
		}
	}

	return timeline, nil
}

func (r RealRepository) AddTweetToUserTimeline(userID int, tweet domain.Tweet) error {
	txn := r.db.Txn(true)
	defer txn.Abort()

	timeline := domain.TimelineTweet{
		// TODO find a better way to generate IDs. Maybe uuids?
		ID:        int(time.Now().UnixMilli()),
		TweetID:   tweet.ID,
		Text:      tweet.Text,
		CreatedAt: tweet.CreatedAt,
		AuthorID:  tweet.AuthorID,
		Follower:  userID,
	}

	if err := txn.Insert("timelineTweets", &timeline); err != nil {
		return err
	}
	txn.Commit()
	return nil
}
