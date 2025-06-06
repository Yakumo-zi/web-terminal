package asset

import (
	"github.com/Yakumo-zi/web-terminal/internal/apiserver/service"
	"github.com/Yakumo-zi/web-terminal/internal/util"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ListRequest struct {
	Offset int `validate:"gte=0",json:"offset"`
	Limit  int `validate:"required,gt=0,lte=100",json:"id",`
}

type ListItem struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
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
	asset, total, err := c.svc.List(ctx.Request().Context(), &service.ListOptions{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	items := make([]ListItem, len(asset))
	for i, item := range asset {
		items[i] = ListItem{
			Id:   item.Id.String(),
			Type: item.Type,
			Name: item.Name,
			Ip:   item.Ip,
			Port: item.Port,
		}
	}
	return ctx.JSON(http.StatusOK, ListResponse{
		Items: items,
		Total: total,
	})

}
