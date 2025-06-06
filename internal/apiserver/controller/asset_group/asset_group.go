package asset_group

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	svc service.AssetGroupService
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc.AssetGroupService,
	}
}

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	controller := NewController(svc)
	e.POST("/Api/V1/AssetGroups", controller.Create)
	e.GET("/Api/V1/AssetGroups/:id", controller.Get)
	e.GET("/Api/V1/AssetGroups", controller.List)
	e.POST("/Api/V1/AssetGroups/:id", controller.Update)
	e.DELETE("/Api/V1/AssetGroups/:id", controller.Delete)
	e.DELETE("/Api/V1/AssetGroups/Collection", controller.DeleteCollection)
	e.POST("/Api/V1/AssetGroups/AddMembers", controller.AddMembers)
}
