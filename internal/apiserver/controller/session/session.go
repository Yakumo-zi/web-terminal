package session

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	svc               service.SessionService
	assetService      service.AssetService
	credentialService service.CredentialService
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc:               svc.SessionService,
		assetService:      svc.AssetService,
		credentialService: svc.CredentialService,
	}
}

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	controller := NewController(svc)
	e.POST("/Api/V1/Sessions", controller.Create)
	e.GET("/Api/V1/Sessions/:id", controller.Get)
	e.GET("/Api/V1/Sessions", controller.List)
	e.POST("/Api/V1/Sessions/:id", controller.Update)
	e.DELETE("/Api/V1/Sessions/:id", controller.Delete)
	e.DELETE("/Api/V1/Sessions/Collection", controller.DeleteCollection)
	e.GET("/Api/V1/Sessions/ByAsset", controller.ByAsset)
}
