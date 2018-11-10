package routers

import "github.com/labstack/echo"

func InitRoutes() *echo.Echo{
	e:= echo.New();
	e = HelloRoute(e)
	e = AuthenticationRoute(e)
	return e
}