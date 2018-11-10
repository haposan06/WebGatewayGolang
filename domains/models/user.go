package models

type User struct {
	ID       int64  `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId	 int64	`json:"role_id"`
	RoleName string `json:"role_name"`
}

func GetDummyUser() *User{
	return &User{ID: 0, Username:"testname", Password:"testpassword"}
}