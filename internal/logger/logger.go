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
    Debug(msg string)
    Info(msg string)
    Warn(msg string)
    Error(msg string)
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
		Level: level,})
    
    logger := SlogLogger {
        logger: slog.New(handler),
    }
    
    return &logger
}

func (t *SlogLogger) Debug(msg string) {
    t.logger.Debug(msg)
}

func (t *SlogLogger) Info(msg string) {
    t.logger.Info(msg)
}

func (t *SlogLogger) Warn(msg string) {
    t.logger.Warn(msg)
}

func (t *SlogLogger) Error(msg string) {
    t.logger.Error(msg)
}
