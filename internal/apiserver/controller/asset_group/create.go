package asset_group

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/domain"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Name string `validate:"required,max=255" json:"name"`
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
	id := uuid.New()
	err := c.svc.Create(ctx.Request().Context(), &domain.AssetGroup{
		Name: req.Name,
		Id:   id,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, CreateResponse{
		Id: id.String(),
	})
}
