package user

import (
	"fmt"
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/utils/db"
	"tc-web-gateway/utils/errors"
)


type DB struct {
	*db.GormDB
}

func NewUserRepository() Repository {
	db := db.GetDBInstance()
	return &DB{db}
}

func (db *DB) FindAll() ([]entities.User, error) {
	var entities []entities.User
	var err error
	err = db.Preload("Role").Find(&entities).Error
	return entities, err
}

func (db *DB) FindById(id int64) (*entities.User, error) {
	var entity entities.User
	var err error
	entity.ID = id
	err = db.Preload("Role").First(&entity, id).Error
	return &entity, err
}

func (db *DB) Store(user *entities.User) (*entities.User, error) {
	var err error
	if db.NewRecord(&user) {
		err = db.Create(&user).Error
	} else {
		err = fmt.Errorf("Save entity failed")
	}
	return user, err
}

func (db *DB) Update(user *entities.User) (*entities.User, error) {
	var err error

	err = db.Save(user).Error
	return user, err
}

func (db *DB) Remove(id int64) error {
	var err error
	var user entities.User
	if id <= 0 {
		return errors.ErrPrimarKey
	}
	user = entities.User{ID:id}
	err = db.Delete(&user).Error
	return err
}

func(db *DB) FindByUsername(username string) (*entities.User, error){
	var entity entities.User
	var err error
	err = db.Where("Username = ?", username).Preload("Role").First(&entity).Error
	return &entity, err
}
