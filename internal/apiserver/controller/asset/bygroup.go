package asset

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ByGroupRequest struct {
	ID     uuid.UUID `query:"id" validate:"required"`
	Offset int       `query:"offset"`
	Limit  int       `query:"limit"`
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

	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	assets, count, err := c.svc.GetByGroup(ctx.Request().Context(), req.ID, req.Offset, req.Limit)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"items": assets,
		"count": count,
	})
}
