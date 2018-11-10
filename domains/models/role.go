package models


type Role struct {
	RoleId int64 	`json:"role_id"`
	RoleName string	`json:"role_name"`
}

func GetDummyRole() *Role{
	return &Role{}
}