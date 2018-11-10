package models

type RolePermission struct {
	RoleId       	int64  		`json:"role_id"`
	PermissionId 	int64		`json:"permission_id"`
	RoleName		string		`json:"role_name"`
	PermissionName	string		`json:"permission_name"`
	Status		 	int64		`json:"status"`
}
