package controller

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/controller/asset"
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"github.com/Yakumo-zi/web-terminal/pkg/web/middlewares"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	apiV1 := e.Group("/Api/V1")
	apiV1.Use(middlewares.LoggerWithSlog(logger.Log()))
	RegisterAssetRoutes(apiV1, svc)
}

func RegisterAssetRoutes(e *echo.Group, svc *service.Service) {
	assetV1 := e.Group("/Assets")
	assetController := asset.NewController(svc)
	{
		assetV1.GET("/", assetController.List)
		assetV1.GET("/:id", assetController.Get)
		assetV1.POST("/", assetController.Create)
		assetV1.POST("/:id", assetController.Update)
		assetV1.DELETE("/:id", assetController.Delete)
		assetV1.DELETE("/Collection", assetController.DeleteCollection)
	}

}
