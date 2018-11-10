package services

import (
	"net/http"
	"tc-web-gateway/domains/models"
	"tc-web-gateway/utils"
)

func Login(user *models.User) (int, string){
	authBackend := utils.InitAuthentication()
	if  valid, err := authBackend.Authenticate(user); err == nil && valid{
		token, err := authBackend.GenerateToken(user)
		if err != nil {
			return http.StatusInternalServerError, ""
		} else {
			return http.StatusOK, token
		}
	}
	return http.StatusUnauthorized, ""
}
