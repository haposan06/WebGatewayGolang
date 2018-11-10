package user

import "tc-web-gateway/domains/entities"

type Repository interface {
	FindAll() ([]entities.User, error)
	FindById(id int64) (*entities.User, error)
	Store(user *entities.User) (*entities.User, error)
	Update(user *entities.User) (*entities.User, error)
	Remove(id int64) error
	FindByUsername(username string) (*entities.User, error)
}

