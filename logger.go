package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger uses the uber zap logger package at its core logger
type Logger struct {
	logger *zap.Logger
	level  zapcore.Level
	output zapcore.WriteSyncer
}

type OptionFunc func(*Logger)
type Level zapcore.Level

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	DPanicLevel
	PanicLevel
	FatalLevel

	DefaultLevel = InfoLevel
)

var logLevels = map[string]Level{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"panic": PanicLevel,
	"fatal": FatalLevel,
}

// New creates a new logger with the given service name
//
// setting log level and custom output are done by
// WithLevel() and WithOutput() funcs of this pkg
// these funcs are completely optional upon calling New()
func New(service string, options ...OptionFunc) *Logger {
	impl := &Logger{
		level:  zapcore.Level(DefaultLevel),
		output: os.Stdout,
	}
	for _, opt := range options {
		opt(impl)
	}
	impl.logger = newJSONLogger(impl.output, Level(impl.level))
	return newZapLogger(impl.logger, service)
}

// logLevel transforms a string log level into zap's log
func logLevel(level string) Level {
	zl, ok := logLevels[strings.ToLower(level)]
	if !ok {
		return DefaultLevel
	}
	return zl
}

func newJSONLogger(output zapcore.WriteSyncer, level Level) *zap.Logger {
	core := zapcore.NewCore(
		newJSONEncoder(), zapcore.Lock(output), zapcore.Level(level))
	return zap.New(core, zap.AddCaller())
}

// newJSONEncoder creates a new JSON log encoder with the default settings.
func newJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        TimeKey,
		LevelKey:       LevelKey,
		NameKey:        NameKey,
		CallerKey:      CallerKey,
		MessageKey:     MessageKey,
		StacktraceKey:  StackTraceKey,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func newZapLogger(logger *zap.Logger, service string) *Logger {
	return &Logger{
		logger: logger.WithOptions(
			zap.AddCallerSkip(1), zap.WithCaller(true)).
			With(
				zap.String("service-name", service)),
	}
}

// Debug logs to the output
//
// if the logging level is set to Debug or above
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

// Info logs to the output
//
// if the logging level is set to Info or above
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs to the output
//
// if the logging level is set to Warn or above
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs to the output
//
// if the logging level is set to Error or above
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

// WithField modifies the logger to add a new field to it
func (l *Logger) WithField(field zap.Field) *Logger {
	return l.WithFields(field)
}

// WithFields returns the logger with the given fields added
func (l *Logger) WithFields(fields ...zap.Field) *Logger {
	if len(fields) == 0 {
		return l
	}
	clone := l.clone()
	clone.logger = l.logger.With(fields...)
	return clone
}

func (l *Logger) clone() *Logger {
	copy := *l
	return &copy
}
