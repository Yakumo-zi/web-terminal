package session

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	AssetId string `validate:"required,uuid" json:"asset_id"`
	CredId  string `validate:"required,uuid" json:"cred_id"`
	Type    string `validate:"required,oneof=ssh rdp" json:"type"`
	Status  string `validate:"required" json:"status"`
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
	credId, err := uuid.Parse(req.CredId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cred, err := c.credentialService.Get(ctx.Request().Context(), credId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	id := uuid.New()
	err = c.svc.Create(ctx.Request().Context(), &domain.Session{
		Type:       req.Type,
		Status:     req.Status,
		Id:         id,
		Asset:      *asset,
		Credential: *cred,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, CreateResponse{
		Id: id.String(),
	})
}
