package logger

import (
    "os"
	"log/slog"
)

type LogLevel int

const (
    Debug LogLevel = iota
	Info
	Warn
	Error
)

type ILogger interface {
    Debug(msg string, keysAndValues ...any)
    Info(msg string, keysAndValues ...any)
    Warn(msg string, keysAndValues ...any)
    Error(msg string, keysAndValues ...any)
}

type SlogLogger struct {
    logger *slog.Logger
}

func New(logLevel LogLevel) ILogger {    

    level := new(slog.LevelVar)
   
    switch logLevel{
    case Debug:
        level.Set(slog.LevelDebug)
    case Info:
        level.Set(slog.LevelInfo)
    case Warn:
        level.Set(slog.LevelWarn)
    case Error:
        level.Set(slog.LevelError)
    }
    
    handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
    })
    
    logger := SlogLogger {
        logger: slog.New(handler),
    }
    
    return &logger
}

func (t *SlogLogger) Debug(msg string, keysAndValues ...any) {

    t.logger.Debug(msg, keysAndValues...)
}

func (t *SlogLogger) Info(msg string, keysAndValues ...any) {
    t.logger.Info(msg, keysAndValues...)
}

func (t *SlogLogger) Warn(msg string, keysAndValues ...any) {
    t.logger.Warn(msg, keysAndValues...)
}

func (t *SlogLogger) Error(msg string, keysAndValues ...any) {
    t.logger.Error(msg, keysAndValues...)
}
