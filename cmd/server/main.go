package main

import (
	"fmt"
	"github.com/Yakumo-zi/web-terminal/api"
	"github.com/Yakumo-zi/web-terminal/service"
	"github.com/labstack/echo/v4"
	"log/slog"
	"strconv"
)

const (
	ApiServerPort = 8000
)

func main() {
	e := echo.New()
	svc := service.NewService()
	api.RegisterRoutes(e, svc)

	svc.BaseLogger.Error("api server listening on 127.0.0.1:"+strconv.Itoa(ApiServerPort),
		slog.Any("error", e.Start(fmt.Sprintf(":%d", ApiServerPort))))
}
