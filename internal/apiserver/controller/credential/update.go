package credential

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	Id       string `validate:"required,uuid" json:"id"`
	Type     string `validate:"required,oneof=password key" json:"type"`
	Username string `validate:"required,max=255" json:"username"`
	Secret   string `validate:"required" json:"secret"`
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
	err = c.svc.Update(ctx.Request().Context(), &domain.Credential{
		Type:     req.Type,
		Username: req.Username,
		Secret:   req.Secret,
		Id:       id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, struct{}{})
}
