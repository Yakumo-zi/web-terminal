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
	e.GET("/Api/V1/AssetGroups/Get", controller.Get)
	e.GET("/Api/V1/AssetGroups/List", controller.List)
	e.POST("/Api/V1/AssetGroups/Create", controller.Create)
	e.POST("/Api/V1/AssetGroups/AddMembers", controller.AddMembers)
	e.POST("/Api/V1/AssetGroups/Update", controller.Update)
	e.POST("/Api/V1/AssetGroups/Delete", controller.Delete)
	e.POST("/Api/V1/AssetGroups/Collection", controller.DeleteCollection)
}
