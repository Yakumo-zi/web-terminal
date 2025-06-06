package asset

import "github.com/Yakumo-zi/web-terminal/internal/apiserver/service"

type Controller struct {
	svc service.AssetService
}

func NewController(svc *service.Service) *Controller {
	return &Controller{
		svc: svc.AssetService,
	}
}
