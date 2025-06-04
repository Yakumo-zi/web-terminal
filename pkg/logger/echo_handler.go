package logger

import (
	"context"
	"github.com/Yakumo-zi/web-terminal/pkg/web/constants"
	"log/slog"
)

type echoSlogHandler struct {
	slog.Handler
}

func (h *echoSlogHandler) Handle(ctx context.Context, r slog.Record) error {
	attr := slog.Group("request",
		slog.String("request_id", ctx.Value(constants.CtxRequestIdKey).(string)),
		slog.String("method", ctx.Value(constants.CtxMethodKey).(string)),
		slog.String("path", ctx.Value(constants.CtxPathKey).(string)),
	)
	r.Add(attr)
	return h.Handler.Handle(ctx, r)
}
