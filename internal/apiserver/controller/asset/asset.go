package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	svc service.AssetService
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc.AssetService,
	}
}

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	controller := NewController(svc)
	e.POST("/Api/V1/Assets/Create", controller.Create)
	e.GET("/Api/V1/Assets/Get", controller.Get)
	e.GET("/Api/V1/Assets/List", controller.List)
	e.POST("/Api/V1/Assets/Update", controller.Update)
	e.POST("/Api/V1/Assets/Delete", controller.Delete)
	e.POST("/Api/V1/Assets/Collection", controller.DeleteCollection)
	e.GET("/Api/V1/Assets/ByGroup", controller.ByGroup)
	e.GET("/Api/V1/Assets/WithoutGroup", controller.WithoutGroup)
}
