package main

import (
	"github.com/labstack/echo"
	"tc-web-gateway/routers"
	"tc-web-gateway/settings"
)

func main(){
	settings.Init()
	e:= echo.New();
	e = routers.InitRoutes()
	e.Logger.Fatal(e.Start(":1323"))
}
