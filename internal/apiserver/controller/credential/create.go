package credential

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	AssetId  string `validate:"required,uuid" json:"asset_id"`
	Type     string `validate:"required,oneof=password key" json:"type"`
	Username string `validate:"required,max=255" json:"username"`
	Secret   string `validate:"required" json:"secret"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

func (c *Controller) Create(ctx echo.Context) error {
	var req CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	if err := validator.Struct(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	assetId, err := uuid.Parse(req.AssetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	asset, err := c.assetService.Get(ctx.Request().Context(), assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id := uuid.New()
	err = c.svc.Create(ctx.Request().Context(), &domain.Credential{
		Type:     req.Type,
		Username: req.Username,
		Secret:   req.Secret,
		Id:       id,
		Asset:    *asset,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, CreateResponse{
		Id: id.String(),
	})
}
