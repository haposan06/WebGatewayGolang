package role

import "tc-web-gateway/domains/entities"

type Repository interface {
	FindAll() ([]entities.Role, error)
	FindById(id int64) (*entities.Role, error)
	Store(role *entities.Role) (*entities.Role, error)
	Update(role *entities.Role) (*entities.Role, error)
	Remove(id int64) error
}