package asset_group

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AddMembersRequest struct {
	GroupID  string   `json:"group_id" validate:"required"`
	AssetIDs []string `json:"asset_ids" validate:"required,min=1"`
}

type AddMembersResponse struct {
	Message string `json:"message"`
}

func (c *Controller) AddMembers(ctx echo.Context) error {
	var req AddMembersRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	groupID, err := uuid.Parse(req.GroupID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	assetIDs := make([]uuid.UUID, len(req.AssetIDs))
	for i, assetID := range req.AssetIDs {
		assetIDs[i], err = uuid.Parse(assetID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}
	}

	err = c.svc.AddMembers(ctx.Request().Context(), groupID, assetIDs)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return ctx.JSON(http.StatusOK, AddMembersResponse{
		Message: "资产已成功添加到资产组",
	})
}
