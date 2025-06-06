package credential

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
	Id       string `json:"id"`
	Type     string `json:"type"`
	Username string `json:"username"`
	Secret   string `json:"secret"`
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
	cid, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cred, err := c.svc.Get(ctx.Request().Context(), cid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, GetResponse{
		Id:       cred.Id.String(),
		Type:     cred.Type,
		Username: cred.Username,
		Secret:   cred.Secret,
	})
}
