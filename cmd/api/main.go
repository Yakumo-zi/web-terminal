package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	os.WriteFile("routes.json", data, 0644)
	e.Start(fmt.Sprintf(":%d", ApiServerPort))
}
