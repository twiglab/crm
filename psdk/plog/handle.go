package plog

import (
	"io"
	"log/slog"
)

func NewHandle(appid string, w io.Writer, level slog.Leveler) slog.Handler {
	return slog.NewJSONHandler(w, &slog.HandlerOptions{AddSource: true, Level: level}).
		WithAttrs([]slog.Attr{slog.String("appid", appid)})
}
