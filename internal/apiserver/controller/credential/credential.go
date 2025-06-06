package credential

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	svc          service.CredentialService
	assetService service.AssetService
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc:          svc.CredentialService,
		assetService: svc.AssetService,
	}
}

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	controller := NewController(svc)
	e.POST("/Api/V1/Credentials", controller.Create)
	e.GET("/Api/V1/Credentials/:id", controller.Get)
	e.GET("/Api/V1/Credentials", controller.List)
	e.POST("/Api/V1/Credentials/:id", controller.Update)
	e.DELETE("/Api/V1/Credentials/:id", controller.Delete)
	e.DELETE("/Api/V1/Credentials/Collection", controller.DeleteCollection)
	e.GET("/Api/V1/Credentials/ByAsset", controller.ByAsset)
}
