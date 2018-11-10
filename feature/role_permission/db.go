package role_permission

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/utils/db"
	"tc-web-gateway/utils/errors"
)
type DB struct {
	*db.GormDB
}

func NewRolePermissionRepository() Repository {
	db := db.GetDBInstance()
	return &DB{db}
}

func (db *DB) FindAll() ([]entities.RolePermission, error) {
	var entities []entities.RolePermission
	var err error
	err = db.Preload("Role").Preload("Permission").Find(&entities).Error
	return entities, err
}

func (db *DB) FindById(role_id, permission_id int64) (*entities.RolePermission, error) {
	var entity entities.RolePermission
	var err error
	err = db.Where("role_id = ? AND permission_Id = ?", role_id, permission_id).Preload("Role").Preload("Permission").First(&entity, role_id, permission_id).Error
	return &entity, err
}

func (db *DB) Store(rolePermission *entities.RolePermission) (*entities.RolePermission, error) {
	var err error
	err = db.Create(&rolePermission).Error
	return rolePermission, err
}

func (db *DB) Update(rolePermission *entities.RolePermission) (*entities.RolePermission, error) {
	var err error

	err = db.Save(&rolePermission).Error
	return rolePermission, err
}

func (db *DB) Remove(role_id, permission_id int64) error {
	var err error
	var rolePermission entities.RolePermission
	if role_id <= 0 || permission_id <= 0{
		return errors.ErrPrimarKey
	}
	rolePermission = entities.RolePermission{RoleId:role_id, PermissionId:permission_id}
	err = db.Delete(&rolePermission).Error
	return err
}
