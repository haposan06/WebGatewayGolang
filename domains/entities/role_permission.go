package entities

import "time"

type RolePermission struct {
	RoleId       	int64  		`gorm:"primary_key"`
	PermissionId 	int64		`gorm:"primary_key"`
	Role			Role
	Permission		Permission
	CreatedBy 		string
	Status	 		int64
	CreatedAt		time.Time
	UpdatedAt		time.Time
}


func (RolePermission) TableName() string {
	return "role_permission"
}