package controller

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/asset"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/asset_group"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/credential"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/session"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"github.com/Yakumo-zi/web-terminal/pkg/web/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	e.Use(middlewares.LoggerWithSlog(logger.Log()))
	asset.RegisterRoutes(e, svc)
	asset_group.RegisterRoutes(e, svc)
	credential.RegisterRoutes(e, svc)
	session.RegisterRoutes(e, svc)
}
