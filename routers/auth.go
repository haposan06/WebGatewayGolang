package routers

import (
	"github.com/labstack/echo"
	"tc-web-gateway/controllers"
)

func AuthenticationRoute(e *echo.Echo) *echo.Echo{
	e.POST("/login", controllers.Login)
	return e
}
