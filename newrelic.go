// Package newrelic maps a birkirb/loggers to a newrelic.Logger
package newrelic

import (
	newrelic "github.com/newrelic/go-agent"
	"gopkg.in/birkirb/loggers.v1"
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	l     loggers.Contextual
	debug bool
}

// NewLogger converts birkirb's logger to a New Relic compatible logger
func NewLogger(logger loggers.Contextual, debug bool) newrelic.Logger {
	return &Logger{
		l:     logger,
		debug: debug,
	}
}

// NewDefaultLogger returns a default Contextual Logger for birkirb's logger.
func NewDefaultLogger() newrelic.Logger {
	l := stdlib.NewDefaultLogger()
	return NewLogger(l, false)
}

// Error logs newrelic errors
func (l Logger) Error(msg string, fields map[string]interface{}) {
	l.l.WithFields(convert(fields)...).Error(msg)
}

// Warn logs newrelic warnings
func (l Logger) Warn(msg string, fields map[string]interface{}) {
	l.l.WithFields(convert(fields)...).Warn(msg)
}

// Info logs newrelic info messages
func (l Logger) Info(msg string, fields map[string]interface{}) {
	l.l.WithFields(convert(fields)...).Info(msg)
}

// Debug logs newrelic debug messages
func (l Logger) Debug(msg string, fields map[string]interface{}) {
	l.l.WithFields(convert(fields)...).Debug(msg)
}

// DebugEnabled returns whether debug is on
func (l Logger) DebugEnabled() bool {
	return l.debug
}

func convert(c map[string]interface{}) []interface{} {
	output := make([]interface{}, 0, 2*len(c))
	for k, v := range c {
		output = append(output, k, v)
	}
	return output
}
