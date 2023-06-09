package logger

import (
	"fmt"
	"io"
	"io/fs"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

const defaultLogFilePermissions fs.FileMode = 0644

// LogrusConfig contains all configurable values for the Logrus logger
type LogrusConfig struct {
	EnableConsole bool
	EnableFile    bool
	Structured    bool
	Level         logrus.Level
	FileLocation  string
}

// LogrusLogger contains all runtime values for using Logrus with the configured output target and input configuration values.
type LogrusLogger struct {
	Config LogrusConfig
	Logger *logrus.Logger
	Output io.Writer
}

// LogrusNestedLogger is a wrapper for Logrus to enable nested logging configuration (loggers that always attach key-value pairs to all log entries)
type LogrusNestedLogger struct {
	Logger *logrus.Entry
}

// NewLogrusLogger creates a new LogrusLogger with the given configuration
func NewLogrusLogger(cfg LogrusConfig) *LogrusLogger {
	appLogger := logrus.New()

	var output io.Writer
	switch {
	case cfg.EnableConsole && cfg.EnableFile:
		logFile, err := os.OpenFile(cfg.FileLocation, os.O_WRONLY|os.O_CREATE, defaultLogFilePermissions)
		if err != nil {
			panic(fmt.Errorf("unable to setup log file: %w", err))
		}
		output = io.MultiWriter(os.Stderr, logFile)
	case cfg.EnableConsole:
		output = os.Stderr
	case cfg.EnableFile:
		logFile, err := os.OpenFile(cfg.FileLocation, os.O_WRONLY|os.O_CREATE, defaultLogFilePermissions)
		if err != nil {
			panic(fmt.Errorf("unable to setup log file: %w", err))
		}
		output = logFile
	default:
		output = io.Discard
	}

	appLogger.SetOutput(output)
	appLogger.SetLevel(cfg.Level)

	if cfg.Structured {
		appLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			PrettyPrint:       false,
		})
	} else {
		appLogger.SetFormatter(&prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			ForceColors:     true,
			ForceFormatting: true,
		})
	}

	return &LogrusLogger{
		Config: cfg,
		Logger: appLogger,
		Output: output,
	}
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func (l *LogrusLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Infof takes a formatted template string and template arguments for the info logging level.
func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func (l *LogrusLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Errorf takes a formatted template string and template arguments for the error logging level.
func (l *LogrusLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Debug logs the given arguments at the debug logging level.
func (l *LogrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

// Info logs the given arguments at the info logging level.
func (l *LogrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

// Warn logs the given arguments at the warning logging level.
func (l *LogrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

// Error logs the given arguments at the error logging level.
func (l *LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

// Debugf takes a formatted template string and template arguments for the debug logging level.
func (l *LogrusNestedLogger) Debugf(format string, args ...interface{}) {
	l.Logger.Debugf(format, args...)
}

// Infof takes a formatted template string and template arguments for the info logging level.
func (l *LogrusNestedLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

// Warnf takes a formatted template string and template arguments for the warning logging level.
func (l *LogrusNestedLogger) Warnf(format string, args ...interface{}) {
	l.Logger.Warnf(format, args...)
}

// Errorf takes a formatted template string and template arguments for the error logging level.
func (l *LogrusNestedLogger) Errorf(format string, args ...interface{}) {
	l.Logger.Errorf(format, args...)
}

// Debug logs the given arguments at the debug logging level.
func (l *LogrusNestedLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

// Info logs the given arguments at the info logging level.
func (l *LogrusNestedLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

// Warn logs the given arguments at the warning logging level.
func (l *LogrusNestedLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

// Error logs the given arguments at the error logging level.
func (l *LogrusNestedLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}
