package asset

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WithoutGroupRequest struct {
	Offset int `query:"offset"`
	Limit  int `query:"limit"`
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

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"items": assets,
		"count": count,
	})
}
