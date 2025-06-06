package main

import (
	"fmt"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
)

const (
	ApiServerPort = 8001
)

func main() {
	e := echo.New()
	svc := service.NewService()
	controller.RegisterRoutes(e, svc)
	e.Start(fmt.Sprintf(":%d", ApiServerPort))
}
