package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/Yakumo-zi/web-terminal/pkg/logger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeleteCollectionRequest struct {
	Ids []string `json:"ids",validate:"required,gt=0,dive,required,uuid"`
}

func (c *Controller) DeleteCollection(ctx echo.Context) error {
	var req DeleteCollectionRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	err := validator.Struct(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	logger.Log().InfoContext(ctx.Request().Context(), "DeleteCollection", "req", req)
	var aids []uuid.UUID
	for _, id := range req.Ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		aids = append(aids, uid)
	}
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err = c.svc.DeleteCollection(ctx.Request().Context(), aids)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, struct{}{})
}
