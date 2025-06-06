package asset

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ByGroupRequest struct {
	ID     string `json:"id" validate:"required"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (c *Controller) ByGroup(ctx echo.Context) error {
	var req ByGroupRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// 设置默认值
	if req.Limit == 0 {
		req.Limit = 20
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	id, err := uuid.Parse(req.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	assets, count, err := c.svc.GetByGroup(ctx.Request().Context(), id, req.Limit, req.Offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	var listResponse ListResponse
	listResponse.Items = make([]ListItem, len(assets))
	for i, asset := range assets {
		listResponse.Items[i] = ListItem{
			Id:   asset.Id.String(),
			Name: asset.Name,
			Ip:   asset.Ip,
			Port: asset.Port,
			Type: asset.Type,
		}
	}
	listResponse.Total = count
	return ctx.JSON(http.StatusOK, listResponse)
}
