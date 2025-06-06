package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type GetRequest struct {
	Id string `param:"id",validate:"required,uuid"`
}

type GetResponse struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

func (c *Controller) Get(ctx echo.Context) error {
	var req GetRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	err := validator.Struct(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	aid, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	asset, err := c.svc.Get(ctx.Request().Context(), aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetResponse{
		Id:   asset.Id.String(),
		Type: asset.Type,
		Name: asset.Name,
		Ip:   asset.Ip,
		Port: asset.Port,
	})
}
