package asset_group

import (
	"net/http"

	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/labstack/echo/v4"
)

type ListRequest struct {
	Offset int `validate:"gte=0" query:"offset"`
	Limit  int `validate:"required,gt=0,lte=100" query:"limit"`
}

type ListItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ListResponse struct {
	Items []ListItem `json:"items"`
	Total int        `json:"total"`
}

func (c *Controller) List(ctx echo.Context) error {
	var req ListRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	validator := util.GetValidator()
	err := validator.Struct(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	groups, total, err := c.svc.List(ctx.Request().Context(), &service.ListOptions{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	items := make([]ListItem, len(groups))
	for i, item := range groups {
		items[i] = ListItem{
			Id:   item.Id.String(),
			Name: item.Name,
		}
	}
	return ctx.JSON(http.StatusOK, ListResponse{
		Items: items,
		Total: total,
	})
}
