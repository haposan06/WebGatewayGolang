package entities

import "time"

type User struct {
	ID       	int64  		`gorm:"primary_key"`
	Username 	string		`gorm:"column:username"`
	Password 	string		`gorm:"column:password"`
	Role		Role
	RoleId		int64
	CreatedBy 	string		`gorm:"column:created_by"`
	CreatedAt	*time.Time	`gorm:"column:created_at"`
	UpdatedAt	*time.Time	`gorm:"column:updated_at"`
	LastLogin	*time.Time	`gorm:"column:last_login"`
}

func GetDummyUser() *User{
	return &User{}
}

func (User) TableName() string {
	return "user"
}