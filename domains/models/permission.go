package models

type Permission struct {
	PermissionID int64 	`json:"permission_id"`
	Name string			`json:"name"`
	Type string			`json:"type"`
	Value string		`json:"value"`
}

func GetDummyPagePermission() *Permission{
	return &Permission{PermissionID: 1, Name:"View Hello Page", Type: "page", Value:"/hello"}
}

func GetDummyAPIPermission() *Permission{
	return &Permission{PermissionID: 2, Name:"API get Hello", Type: "api", Value:"/gethello"}
}