package api

import (
	"fmt"
	"github.com/Yakumo-zi/web-terminal/internal/web/middlewares"
	"github.com/Yakumo-zi/web-terminal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterRoutes(e *echo.Echo, svc *service.Service) {
	apiV1 := e.Group("/api/v1")
	apiV1.Use(middlewares.LoggerWithSlog(svc.WebLogger))
	apiV1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	userV1 := apiV1.Group("/user")
	userV1.GET("/:id", func(c echo.Context) error {
		id := c.Param("id")
		return c.String(http.StatusOK, id)
	})
	errorV1 := apiV1.Group("/error")
	errorV1.GET("/:code", func(c echo.Context) error {
		code := c.Param("code")
		c.Response().WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("it's a error code: %s", code)
	})
}
