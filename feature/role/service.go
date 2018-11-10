package role

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/domains/models"
)

type Service interface {
	FindAll() ([]models.Role, error)
	FindById(id int64) (*models.Role, error)
	Store(role *models.Role) (*models.Role, error)
	Update(role *models.Role) (*models.Role, error)
	Remove(id int64) error
}


type RoleService struct {
	Repository Repository``
}

func NewRoleService() Service {
	return &RoleService{Repository:NewRoleRepository()}
}

func (u *RoleService) FindAll() ([]models.Role, error) {
	var roles []models.Role
	entities, err := u.Repository.FindAll()
	roles = make([]models.Role, len(entities))
	if err == nil {
		for i, v := range entities {
			roles[i].RoleId = v.ID
			roles[i].RoleName = v.RoleName
		}
	}
	return roles, err
}

func (u *RoleService) FindById(id int64) (*models.Role, error) {
	var role models.Role = *new(models.Role)
	entity, err := u.Repository.FindById(id)
	if err == nil {
		role.RoleId = entity.ID
		role.RoleName = entity.RoleName
	}
	return &role, err
}

func (u *RoleService) Store(role *models.Role) (*models.Role, error) {
	var entity *entities.Role = new(entities.Role)
	entity.ID = role.RoleId
	entity.RoleName = role.RoleName
	entity, err := u.Repository.Store(entity)
	if err == nil {
		role.RoleId = entity.ID
	}
	return role,err
}

func (u *RoleService) Update(role *models.Role) (*models.Role, error) {
	var entity *entities.Role = new(entities.Role)
	entity = new(entities.Role)
	entity.ID = role.RoleId
	entity.RoleName = role.RoleName
	entity, err := u.Repository.Update(entity)
	return role,err
}

func (u *RoleService) Remove(id int64) error {
	err := u.Repository.Remove(id)
	return err
}