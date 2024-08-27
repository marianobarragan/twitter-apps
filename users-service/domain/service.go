package domain

import (
	"log"
)

type Service interface {
	GetUser(id int) (User, bool, error)
	SaveUser(user User) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return service{
		repository: r,
	}
}

func (ps service) GetUser(id int) (User, bool, error) {
	return ps.repository.GetUser(id)
}

func (ps service) SaveUser(user User) (User, error) {
	user, err := ps.repository.PostUser(user)
	log.Println("Created user ID")
	log.Println(user.ID)
	return user, err
}
