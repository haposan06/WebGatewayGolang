package entities

import "time"

type Permission struct {
	ID       		int64  		`gorm:"primary_key"`
	PermissionName 	string
	Type 			string
	Value		 	string
	CreatedBy 		string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

func GetDummyPermission() *Permission{
	return &Permission{}
}

func (Permission) TableName() string {
	return "permission"
}