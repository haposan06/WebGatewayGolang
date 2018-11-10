package permission

import "tc-web-gateway/domains/entities"

type Repository interface {
	FindAll() ([]entities.Permission, error)
	FindById(id int64) (*entities.Permission, error)
	Store(permission *entities.Permission) (*entities.Permission, error)
	Update(permission *entities.Permission) (*entities.Permission, error)
	Remove(id int64) error
}