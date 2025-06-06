package logger

import (
	"io"
	"log/slog"
	"os"
	"sync"
)

var webLogger *slog.Logger
var webOnce sync.Once

func Log() *slog.Logger {
	webOnce.Do(func() {
		webLogger = NewWebLogger(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug})
	})
	return webLogger
}

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
