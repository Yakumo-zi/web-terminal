package session

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ByAssetRequest struct {
	ID     string `query:"id" validate:"required"`
	Offset int    `query:"offset"`
	Limit  int    `query:"limit"`
}

func (c *Controller) ByAsset(ctx echo.Context) error {
	var req ByAssetRequest
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

	sessions, count, err := c.svc.GetByAsset(ctx.Request().Context(), id, req.Limit, req.Offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"items": sessions,
		"count": count,
	})
}
