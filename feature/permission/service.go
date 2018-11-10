package permission

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/domains/models"
)

type Service interface {
	FindAll() ([]models.Permission, error)
	FindById(id int64) (*models.Permission, error)
	Store(permission *models.Permission) (*models.Permission, error)
	Update(permission *models.Permission) (*models.Permission, error)
	Remove(id int64) error
}


type PermissionService struct {
	Repository Repository``
}

func NewPermissionService() Service {
	return &PermissionService{Repository:NewPermissionRepository()}
}

func (u *PermissionService) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	entities, err := u.Repository.FindAll()
	permissions = make([]models.Permission, len(entities))
	if err == nil {
		for i, v := range entities {
			permissions[i].PermissionID = v.ID
			permissions[i].Name = v.PermissionName
			permissions[i].Type = v.Type
			permissions[i].Value = v.Value
		}
	}
	return permissions, err
}

func (u *PermissionService) FindById(id int64) (*models.Permission, error) {
	var permission models.Permission = *new(models.Permission)
	entity, err := u.Repository.FindById(id)
	if err == nil {
		permission.PermissionID = entity.ID
		permission.Name = entity.PermissionName
		permission.Type = entity.Type
		permission.Value = entity.Value
	}
	return &permission, err
}

func (u *PermissionService) Store(permission *models.Permission) (*models.Permission, error) {
	var entity *entities.Permission = new(entities.Permission)
	entity.ID = permission.PermissionID
	entity.PermissionName = permission.Name
	entity.Type = permission.Type
	entity.Value = permission.Value
	entity, err := u.Repository.Store(entity)
	if err == nil {
		permission.PermissionID = entity.ID
	}
	return permission,err
}

func (u *PermissionService) Update(permission *models.Permission) (*models.Permission, error) {
	var entity *entities.Permission = new(entities.Permission)
	entity.ID = permission.PermissionID
	entity.PermissionName = permission.Name
	entity.Type = permission.Type
	entity.Value = permission.Value
	entity, err := u.Repository.Update(entity)
	return permission,err
}

func (u *PermissionService) Remove(id int64) error {
	err := u.Repository.Remove(id)
	return err
}