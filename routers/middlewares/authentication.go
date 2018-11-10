package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"tc-web-gateway/utils"
)

func RequireTokenAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var tokenString string
		var err error
		auth := utils.InitAuthentication()
		if tokenString, err = getTokenFromHeader(c.Request()) ; err != nil{
			c.JSON(http.StatusUnauthorized, map[string]string{
				"status": "Malformed header",
			})
			return nil
		}
		token, err := auth.CheckTokenValidity(tokenString)
		if err != nil {
			return nil
		}
		claims:= token.Claims.(jwt.MapClaims)
		role := claims["role"]
		authorize:= utils.Init()
		status, err := authorize.CheckAccessRules(role.(string), c.Request())
		if status == http.StatusOK{
			next(c)
		}
		return nil
	}
}

func getTokenFromHeader(req *http.Request) (string, error) {
	var tokenString string
	var err error

	tokenRaw := req.Header.Get("Authorization")
	if len(tokenRaw) <= 0 {
		return "", fmt.Errorf("There is no header Authorization")
	}
	tokenRawString := strings.Split(tokenRaw, "Bearer ")
	tokenString = tokenRawString[1]

	if len(tokenString) <= 0 {
		return "", fmt.Errorf("Token not found")
	}
	return tokenString, err
}

