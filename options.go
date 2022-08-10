package log

import "go.uber.org/zap/zapcore"

// WithLevel allows setting a custom level to the logger
func WithLevel(level string) OptionFunc {
	return func(log *Logger) {
		log.level = zapcore.Level(logLevel(level))
	}
}

// WithOutput allows setting a custom output to the logger
func WithOutput(output zapcore.WriteSyncer) OptionFunc {
	return func(log *Logger) {
		log.output = output
	}
}
