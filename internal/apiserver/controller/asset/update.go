package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UpdateRequest struct {
	Id   string `json:"id",validate:"required,uuid"`
	Name string `json:"name",validate:"required,max=255"`
	Ip   string `json:"ip",validate:"required,ip"`
	Port int    `json:"port",validate:"required,min=1,max=65535"`
	Type string `json:"type",validate:"required,oneof=host db"`
}

func (c *Controller) Update(ctx echo.Context) error {
	var req UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	if err := validator.Struct(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.svc.Update(ctx.Request().Context(), &domain.Asset{
		Name: req.Name,
		Ip:   req.Ip,
		Port: req.Port,
		Type: req.Type,
		Id:   id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, struct{}{})
}
