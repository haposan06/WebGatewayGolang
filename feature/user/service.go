package user

import (
	"tc-web-gateway/domains/entities"
	"tc-web-gateway/domains/models"
	"tc-web-gateway/utils/tools"
)

type Service interface {
	FindAll() ([]models.User, error)
	FindById(id int64) (*models.User, error)
	Store(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Remove(id int64) error
	FindByUsername(username string) (*models.User, error)
}

type UserService struct {
	Repository Repository``
}

func NewUserService() Service {
	return &UserService{Repository:NewUserRepository()}
}

func (u *UserService) FindAll() ([]models.User, error) {
	var users []models.User
	entities, err := u.Repository.FindAll()
	users = make([]models.User, len(entities))
	if err == nil {
		for i, v := range entities {
			users[i].ID = v.ID
			users[i].Username = v.Username
			users[i].RoleName = v.Role.RoleName
		}
	}
	return users, err
}

func (u *UserService) FindById(id int64) (*models.User, error) {
	var user models.User = *new(models.User)
	entity, err := u.Repository.FindById(id)
	if err == nil {
		user.ID = entity.ID
		user.Username = entity.Username
		user.RoleName = entity.Role.RoleName
	}
	return &user, err
}

func (u *UserService) Store(user *models.User) (*models.User, error) {
	var entity *entities.User = new(entities.User)
	hashPassword, err := tools.GenerateHashPassword(user.Password)
	entity.Username = user.Username
	entity.Password = hashPassword
	entity.RoleId = user.RoleId
	//TODO implement last_login
	//*entity.LastLogin = time.Now()
	entity, err = u.Repository.Store(entity)
	if err == nil {
		user.ID = entity.ID
	}
	return user,err
}

func (u *UserService) Update(user *models.User) (*models.User, error) {
	var entity *entities.User = new(entities.User)
	hashPassword, err := tools.GenerateHashPassword(user.Password)
	entity = new(entities.User)
	entity.ID = user.ID
	entity.Username = user.Username
	entity.Password = hashPassword
	entity.RoleId = user.RoleId
	entity, err = u.Repository.Update(entity)
	return user,err
}

func (u *UserService) Remove(id int64) error {
	err := u.Repository.Remove(id)
	return err
}

func (u *UserService) FindByUsername(username string) (*models.User, error) {
	var user models.User = *new(models.User)
	entity, err := u.Repository.FindByUsername(username)
	if err == nil {
		user.ID = entity.ID
		user.Username = entity.Username
		user.Password = entity.Password
		user.RoleName = entity.Role.RoleName
	}
	return &user, err
}