package main

import (
	"fmt"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
	"log/slog"
	"strconv"
)

const (
	ApiServerPort = 8001
)

func main() {
	e := echo.New()
	svc := service.NewService()
	controller.RegisterRoutes(e, svc)

	svc.BaseLogger.Error("api api listening on 127.0.0.1:"+strconv.Itoa(ApiServerPort),
		slog.Any("error", e.Start(fmt.Sprintf(":%d", ApiServerPort))))
}
