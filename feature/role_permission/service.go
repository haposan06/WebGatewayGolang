package role_permission

import (
	"github.com/labstack/echo"
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/domains/models"
	"tc-web-gateway/utils"
)

type Service interface {
	FindAll() ([]models.RolePermission, error)
	FindById(role_id, permission_id int64) (*models.RolePermission, error)
	Store(rolePermission *models.RolePermission) (*models.RolePermission, error)
	Update(rolePermission *models.RolePermission) (*models.RolePermission, error)
	Remove(role_id int64,permission_id int64) error
}

type RolePermissionService struct {
	Repository Repository``
}

func NewRolePermissionService() Service {
	return &RolePermissionService{Repository:NewRolePermissionRepository()}
}

func (u *RolePermissionService) FindAll() ([]models.RolePermission, error) {
	var rolePermissions []models.RolePermission
	entities, err := u.Repository.FindAll()
	rolePermissions = make([]models.RolePermission, len(entities))
	if err == nil {
		for i, v := range entities {
			rolePermissions[i].RoleId = v.RoleId
			rolePermissions[i].RoleName = v.Role.RoleName
			rolePermissions[i].PermissionId = v.PermissionId
			rolePermissions[i].PermissionName = v.Permission.PermissionName
			rolePermissions[i].Status = v.Status
		}
	}
	return rolePermissions, err
}

func (u *RolePermissionService) FindById(role_id, permission_id int64) (*models.RolePermission, error) {
	var rolePermission models.RolePermission =  *new(models.RolePermission)
	entity, err := u.Repository.FindById(role_id, permission_id)
	if err == nil {
		rolePermission.RoleId = entity.RoleId
		rolePermission.RoleName = entity.Role.RoleName
		rolePermission.PermissionId = entity.PermissionId
		rolePermission.PermissionName = entity.Permission.PermissionName
		rolePermission.Status = entity.Status
	}
	return &rolePermission, err
}

func (u *RolePermissionService) Store(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	var entity *entities.RolePermission = new(entities.RolePermission)
	entity.RoleId = rolePermission.RoleId
	entity.PermissionId = rolePermission.PermissionId
	entity.Status = rolePermission.Status
	entity, err := u.Repository.Store(entity)
	if err ==nil{
		entity, err = u.Repository.FindById(rolePermission.RoleId, rolePermission.PermissionId)
		utils.AddPolicy(entity.Role.RoleName, entity.Permission.Value, "*")
	}
	return rolePermission,err
}

func (u *RolePermissionService) Update(rolePermission *models.RolePermission) (*models.RolePermission, error) {
	var entity *entities.RolePermission = new(entities.RolePermission)
	entity.RoleId = rolePermission.RoleId
	entity.PermissionId = rolePermission.PermissionId
	entity.Status = rolePermission.Status
	entity, err := u.Repository.Update(entity)
	return rolePermission,err
}

func (u *RolePermissionService) Remove(role_id, permission_id int64) error {
	rolePermission, err:= u.Repository.FindById(role_id, permission_id)
	if err != nil {
		return echo.ErrNotFound
	}
	err = u.Repository.Remove(role_id, permission_id)
	if err == nil {
		utils.RemovePolicy(rolePermission.Role.RoleName, rolePermission.Permission.Value, "*")
	}
	return err
}