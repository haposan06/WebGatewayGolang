package role

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/utils/db"
	"tc-web-gateway/utils/errors"
)

type DB struct {
	*db.GormDB
}

func NewRoleRepository() Repository {
	db := db.GetDBInstance()
	return &DB{db}
}

func (db *DB) FindAll() ([]entities.Role, error) {
	var entities []entities.Role
	var err error
	err = db.Find(&entities).Error
	return entities, err
}

func (db *DB) FindById(id int64) (*entities.Role, error) {
	var entity entities.Role
	var err error
	entity.ID = id
	err = db.First(&entity, id).Error
	return &entity, err
}

func (db *DB) Store(role *entities.Role) (*entities.Role, error) {
	var err error
	if db.NewRecord(&role) {
		err = db.Create(&role).Error
	} else {
		err = errors.ErrDBOpsFailed
	}
	return role, err
}

func (db *DB) Update(role *entities.Role) (*entities.Role, error) {
	var err error

	err = db.Save(role).Error
	return role, err
}

func (db *DB) Remove(id int64) error {
	var err error
	var role entities.Role
	if id <= 0 {
		return errors.ErrPrimarKey
	}
	role = entities.Role{ID:id}
	err = db.Delete(&role).Error
	return err
}
