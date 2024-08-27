package gomemdb

import (
	"errors"
	"github.com/hashicorp/go-memdb"
	"time"
	"users-service/domain"
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
			"user": &memdb.TableSchema{
				Name: "user",
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

func (r RealRepository) GetUser(id int) (domain.User, bool, error) {
	txn := r.db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "id", id)
	if err != nil {
		return domain.User{}, false, err
	}

	if raw == nil {
		return domain.User{}, false, nil
	}

	user, ok := raw.(*domain.User)
	if !ok {
		return domain.User{}, true, errors.New("unable to convert row to struct")
	}

	return *user, true, nil
}

func (r RealRepository) PostUser(user domain.User) (domain.User, error) {
	txn := r.db.Txn(true)
	defer txn.Abort()
	// TODO find a better way to generate IDs. Maybe uuids?
	user.ID = int(time.Now().UnixMilli())
	user.CreatedAt = time.Now()
	if err := txn.Insert("user", &user); err != nil {
		return user, err
	}
	txn.Commit()
	return user, nil
}
