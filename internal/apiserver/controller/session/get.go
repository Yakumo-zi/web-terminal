package session

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetRequest struct {
	Id string `validate:"required,uuid" param:"id"`
}

type GetResponse struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
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
	sid, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	sess, err := c.svc.Get(ctx.Request().Context(), sid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetResponse{
		Id:     sess.Id.String(),
		Type:   sess.Type,
		Status: sess.Status,
	})
}
