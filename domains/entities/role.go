package entities

import "time"

type Role struct {
	ID       	int64  		`gorm:"primary_key"`
	RoleName 	string		`gorm:"column:role_name"`
	CreatedBy 	string		`gorm:"created_by"`
	CreatedAt	time.Time	`gorm:"column:created_at"`
	UpdatedAt	time.Time	`gorm:"column:updated_at"`
}

func GetDummyRole() *Role{
	return &Role{}
}

func (Role) TableName() string {
	return "role"
}