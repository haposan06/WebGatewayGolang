package routers

import (
	"github.com/labstack/echo"
	"tc-web-gateway/controllers"
	"tc-web-gateway/routers/middlewares"
)

func HelloRoute(e *echo.Echo) *echo.Echo {
	e.POST("/hore", controllers.HelloWorld , middlewares.RequireTokenAuthentication )

	return e;
}
