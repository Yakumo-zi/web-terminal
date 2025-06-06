package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CreateRequest struct {
	Name string `validate:"required,max=255",json:"name"`
	Ip   string `validate:"required,ip",json:"ip"`
	Port int    `validate:"required,min=1,max=65535",json:"port"`
	Type string `validate:"required,oneof=host db",json:"type"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func (c *Controller) Create(ctx echo.Context) error {
	var req CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := uuid.New()
	err := c.svc.Create(ctx.Request().Context(), &domain.Asset{
		Name: req.Name,
		Ip:   req.Ip,
		Port: req.Port,
		Type: req.Type,
		Id:   id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, CreateResponse{
		Id: id.String(),
	})
}
