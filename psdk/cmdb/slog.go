package cmdb

import (
	"io"
	"log/slog"
	"strings"

	"github.com/twiglab/crm/psdk/plog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const DefaultLogFile = "logs/log.log"

func NewRollingWriter(logFile string) io.Writer {
	if logFile == "" {
		return &lumberjack.Logger{
			Filename: DefaultLogFile,
		}
	}
	return &lumberjack.Logger{
		Filename: logFile,
	}
}

func RootLogger(app App) (*slog.Logger, *slog.LevelVar) {
	level := &slog.LevelVar{}
	level.Set(parse(app.Debug))

	w := NewRollingWriter(app.LogFile)
	h := plog.NewHandle(app.ID, w, level)

	rlog := slog.New(h)
	slog.SetDefault(rlog)

	return rlog, level
}

func parse(s string) slog.Level {
	if s == "" {
		return slog.LevelDebug
	}
	switch strings.ToUpper(s) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}
