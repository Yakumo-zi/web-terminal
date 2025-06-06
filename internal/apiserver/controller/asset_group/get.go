package asset_group

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetRequest struct {
	Id string `validate:"required,uuid" json:"id"`
}

type GetResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
	group, err := c.svc.Get(ctx.Request().Context(), aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetResponse{
		Id:   group.Id.String(),
		Name: group.Name,
	})
}
