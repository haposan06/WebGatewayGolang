package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"tc-web-gateway/domains/models"
	"tc-web-gateway/feature/role_permission"
)

func HelloWorld(c echo.Context) error {
	requestUser := new(models.RolePermission)
	decoder := json.NewDecoder(c.Request().Body)
	err := decoder.Decode(requestUser)
	fmt.Println(err)
	service := role_permission.NewRolePermissionService()
	a:= service.Remove(1,4)
	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, a)
}
