package logger

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type Logger struct {
	*slog.Logger
}

type LogLevel string

const (
	LogLevelNone  LogLevel = "none"
	LogLevelError LogLevel = "error"
	LogLevelWarn  LogLevel = "warn"
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
)

var (
	instance *Logger
	once     sync.Once
)

func Init(logFile string, logLevel LogLevel) error {
	var initError error

	once.Do(func() {
		var handler slog.Handler

		if logLevel == LogLevelNone {
			handler = slog.DiscardHandler
		} else {
			var level slog.Level

			switch logLevel {
			case LogLevelError:
				level = slog.LevelError
			case LogLevelWarn:
				level = slog.LevelWarn
			case LogLevelInfo:
				level = slog.LevelInfo
			case LogLevelDebug:
				level = slog.LevelDebug
			}

			opts := &slog.HandlerOptions{
				Level: level,
			}

			const FilePermissionsReadWriteAll os.FileMode = 0o666

			file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, FilePermissionsReadWriteAll)
			if err != nil {
				initError = fmt.Errorf("failed to open log file: %w", err)

				return
			}

			handler = slog.NewJSONHandler(file, opts)
		}

		instance = &Logger{slog.New(handler)}
	})

	return initError
}

func GetInstance() *Logger {
	return instance
}
