package asset_group

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	Id   string `validate:"required,uuid" param:"id"`
	Name string `validate:"required,max=255" json:"name"`
}

func (c *Controller) Update(ctx echo.Context) error {
	var req UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.svc.Update(ctx.Request().Context(), &domain.AssetGroup{
		Name: req.Name,
		Id:   id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, struct{}{})
}
