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
	apiV1 := e.Group("/Api/V1")
	apiV1.Use(middlewares.LoggerWithSlog(logger.Log()))
	RegisterAssetRoutes(apiV1, svc)
	RegisterAssetGroupRoutes(apiV1, svc)
	RegisterCredentialRoutes(apiV1, svc)
	RegisterSessionRoutes(apiV1, svc)
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

func RegisterAssetGroupRoutes(e *echo.Group, svc *service.Service) {
	agV1 := e.Group("/AssetGroups")
	agController := asset_group.NewController(svc)
	{
		agV1.GET("/", agController.List)
		agV1.GET("/:id", agController.Get)
		agV1.POST("/", agController.Create)
		agV1.POST("/:id", agController.Update)
		agV1.DELETE("/:id", agController.Delete)
		agV1.DELETE("/Collection", agController.DeleteCollection)
	}
}

func RegisterCredentialRoutes(e *echo.Group, svc *service.Service) {
	credV1 := e.Group("/Credentials")
	credController := credential.NewController(svc)
	{
		credV1.GET("/", credController.List)
		credV1.GET("/:id", credController.Get)
		credV1.POST("/", credController.Create)
		credV1.POST("/:id", credController.Update)
		credV1.DELETE("/:id", credController.Delete)
		credV1.DELETE("/Collection", credController.DeleteCollection)
	}
}

func RegisterSessionRoutes(e *echo.Group, svc *service.Service) {
	sessV1 := e.Group("/Sessions")
	sessController := session.NewController(svc)
	{
		sessV1.GET("/", sessController.List)
		sessV1.GET("/:id", sessController.Get)
		sessV1.POST("/", sessController.Create)
		sessV1.POST("/:id", sessController.Update)
		sessV1.DELETE("/:id", sessController.Delete)
		sessV1.DELETE("/Collection", sessController.DeleteCollection)
	}
}
