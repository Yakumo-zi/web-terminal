package asset

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name        string `validate:"required,max=255" json:"name"`
	Ip          string `validate:"required,ip" json:"ip"`
	Port        int    `validate:"required,min=1,max=65535" json:"port"`
	Type        string `validate:"required,oneof=host db" json:"type"`
	Description string `json:"description"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func (c *Controller) Create(ctx echo.Context) error {
	var req CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	id := uuid.New()
	asset := &domain.Asset{
		Id:   id,
		Name: req.Name,
		Ip:   req.Ip,
		Port: req.Port,
		Type: req.Type,
	}
	err := c.svc.Create(ctx.Request().Context(), asset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return ctx.JSON(http.StatusOK, CreateResponse{
		Id: id.String(),
	})
}
