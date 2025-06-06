package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WithoutGroupRequest struct {
	Offset int `query:"offset" validate:"gte=0"`
	Limit  int `query:"limit"  validate:"required,gt=0,lte=100"`
}

func (c *Controller) WithoutGroup(ctx echo.Context) error {
	var req WithoutGroupRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	// 设置默认值
	if req.Limit == 0 {
		req.Limit = 20
	}

	assets, count, err := c.svc.GetWithoutGroup(ctx.Request().Context(), req.Limit, req.Offset)
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
