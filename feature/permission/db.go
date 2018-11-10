package permission

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/utils/db"
	"tc-web-gateway/utils/errors"
)

type DB struct {
	*db.GormDB
}

func NewPermissionRepository() Repository {
	db := db.GetDBInstance()
	return &DB{db}
}

func (db *DB) FindAll() ([]entities.Permission, error) {
	var entities []entities.Permission
	var err error
	err = db.Find(&entities).Error
	return entities, err
}

func (db *DB) FindById(id int64) (*entities.Permission, error) {
	var entity entities.Permission
	var err error
	entity.ID = id
	err = db.First(&entity, id).Error
	return &entity, err
}

func (db *DB) Store(permission *entities.Permission) (*entities.Permission, error) {
	var err error
	if db.NewRecord(&permission) {
		err = db.Create(&permission).Error
	} else {
		err = errors.ErrDBOpsFailed
	}
	return permission, err
}

func (db *DB) Update(permission *entities.Permission) (*entities.Permission, error) {
	var err error

	err = db.Save(permission).Error
	return permission, err
}

func (db *DB) Remove(id int64) error {
	var err error
	var permission entities.Permission
	if id <= 0 {
		return errors.ErrPrimarKey
	}
	permission = entities.Permission{ID:id}
	err = db.Delete(&permission).Error
	return err
}
