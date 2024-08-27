package domain

type Repository interface {
	GetUser(id int) (User, bool, error)
	PostUser(user User) (User, error)
}
