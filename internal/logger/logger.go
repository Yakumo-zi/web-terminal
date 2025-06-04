package logger

import (
	"io"
	"log/slog"
)

func NewLogger(w io.Writer) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(w, opts))
	return logger
}

func NewWebLogger(w io.Writer, opts *slog.HandlerOptions) *slog.Logger {
	logger := slog.New(&echoSlogHandler{slog.NewJSONHandler(w, opts)})
	return logger
}
