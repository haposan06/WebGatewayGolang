package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"tc-web-gateway/domains/models"
	"tc-web-gateway/services"
)

func Login(c echo.Context) error {
	requestUser := new(models.User)
	decoder := json.NewDecoder(c.Request().Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	tokenRes := &models.TokenAuthentication{token}
	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(responseStatus, tokenRes)
}
