package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeletetRequest struct {
	Id string `param:"id",validate:"required,uuid"`
}

func (c *Controller) Delete(ctx echo.Context) error {
	var req DeletetRequest
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
	err = c.svc.Delete(ctx.Request().Context(), aid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, struct{}{})
}
