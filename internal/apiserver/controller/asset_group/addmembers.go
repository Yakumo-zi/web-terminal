package asset_group

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AddMembersRequest struct {
	GroupID  uuid.UUID   `json:"group_id" validate:"required"`
	AssetIDs []uuid.UUID `json:"asset_ids" validate:"required,min=1"`
}

type AddMembersResponse struct {
	Message string `json:"message"`
}

func (c *Controller) AddMembers(ctx echo.Context) error {
	var req AddMembersRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	if err := ctx.Validate(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err := c.svc.AddMembers(ctx.Request().Context(), req.GroupID, req.AssetIDs)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, AddMembersResponse{
		Message: "资产已成功添加到资产组",
	})
}
