package middlewares

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/internal/web/constants"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

func LoggerWithSlog(logger *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			uuid, err := uuid.NewUUID()
			if err != nil {
				return err
			}
			ctx = context.WithValue(ctx, constants.CtxRequestIdKey, uuid.String())
			ctx = context.WithValue(ctx, constants.CtxMethodKey, c.Request().Method)
			ctx = context.WithValue(ctx, constants.CtxPathKey, c.Request().URL.Path)
			start := time.Now()
			err = next(c)
			if err != nil {
				end := time.Now()
				logger.ErrorContext(ctx, "api request error",
					slog.Int("status_code", c.Response().Status),
					slog.Any("error", err),
					slog.Int64("latency", int64(end.Sub(start).Milliseconds())),
				)
				c.Error(err)
				return err
			}
			end := time.Now()
			logger.InfoContext(ctx, "api request success",
				slog.Int("status_code", c.Response().Status),
				slog.Int64("latency", int64(end.Sub(start).Milliseconds())),
			)
			return nil
		}
	}
}
