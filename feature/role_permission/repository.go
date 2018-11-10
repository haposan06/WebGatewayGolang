package role_permission

import "tc-web-gateway/domains/entities"

type Repository interface {
	FindAll() ([]entities.RolePermission, error)
	FindById(role_id int64, permission_id int64) (*entities.RolePermission, error)
	Store(user *entities.RolePermission) (*entities.RolePermission, error)
	Update(user *entities.RolePermission) (*entities.RolePermission, error)
	Remove(role_id, permission_id int64) error
}