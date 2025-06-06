package logger

import (
	"context"
	"log/slog"

	"github.com/Yakumo-zi/web-terminal/pkg/web/constants"
)

type echoSlogHandler struct {
	slog.Handler
}

func (h *echoSlogHandler) Handle(ctx context.Context, r slog.Record) error {
	if ctx.Value(constants.CtxRequestIdKey) == nil {
		return h.Handler.Handle(ctx, r)
	}
	request_id := ctx.Value(constants.CtxRequestIdKey)
	method := ctx.Value(constants.CtxMethodKey)
	path := ctx.Value(constants.CtxPathKey)
	attr := slog.Group("request",
		slog.String("request_id", request_id.(string)),
		slog.String("method", method.(string)),
		slog.String("path", path.(string)),
	)
	r.Add(attr)
	return h.Handler.Handle(ctx, r)
}
