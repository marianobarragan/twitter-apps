package gomemdb

import (
	"errors"
	"time"
	"tweets-service/domain"

	"github.com/hashicorp/go-memdb"
)

type RealRepository struct {
	db *memdb.MemDB
}

var _ domain.Repository = RealRepository{}
var _ domain.Transaction = &memdb.Txn{}

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
			"tweet": &memdb.TableSchema{
				Name: "tweet",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.IntFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
}

func (r RealRepository) GetTweet(id int) (domain.Tweet, bool, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("tweet", "id", id)
	if err != nil {
		return domain.Tweet{}, false, err
	}

	if raw == nil {
		return domain.Tweet{}, false, nil
	}

	tweet, ok := raw.(*domain.Tweet)
	if !ok {
		return domain.Tweet{}, true, errors.New("unable to convert row to struct")
	}

	return *tweet, true, nil
}

func (r RealRepository) PostTweet(tweet domain.Tweet) (domain.Tweet, domain.Transaction, error) {
	txn := r.db.Txn(true)
	// TODO find a better way to generate IDs. Maybe uuids?
	tweet.ID = int(time.Now().UnixMilli())
	tweet.CreatedAt = time.Now()
	if err := txn.Insert("tweet", &tweet); err != nil {
		txn.Abort()
		return tweet, nil, err
	}
	return tweet, txn, nil
}

func (r RealRepository) SearchTweets() ([]domain.Tweet, error) {
	//TODO implement me
	panic("implement me")
}
