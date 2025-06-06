package middlewares

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/pkg/web/constants"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

func LoggerWithSlog(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := c.Request().Header.Get("X-Request-ID")
			ctx := c.Request().Context()
			if len(requestId) == 0 {
				requestId = uuid.NewString()
			}
			ctx = context.WithValue(ctx, constants.CtxRequestIdKey, requestId)
			ctx = context.WithValue(ctx, constants.CtxMethodKey, c.Request().Method)
			ctx = context.WithValue(ctx, constants.CtxPathKey, c.Request().URL.Path)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
			start := time.Now()
			err := next(c)
			if err != nil {
				end := time.Now()
				logger.ErrorContext(ctx, "api request error",
					slog.Int("status_code", c.Response().Status),
					slog.Any("error", err),
					slog.Int64("latency", end.Sub(start).Milliseconds()),
				)
				c.Error(err)
				return err
			}
			end := time.Now()
			logger.InfoContext(ctx, "api request success",
				slog.Int("status_code", c.Response().Status),
				slog.Int64("latency", end.Sub(start).Milliseconds()),
			)
			return nil
		}
	}
}
