// Package newrelic maps a birkirb/loggers to a newrelic.Logger
package newrelic

import (
	"fmt"
	newrelic "github.com/newrelic/go-agent"
	"gopkg.in/birkirb/loggers.v1"
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
	"regexp"
)

// Logger is an Contextual logger wrapper over Logrus's logger.
type Logger struct {
	l     loggers.Contextual
	debug bool
	censor bool
}

// NewLogger converts birkirb's logger to a New Relic compatible logger. If censor, any new relic license is redacted from logs
func NewLogger(logger loggers.Contextual, debug, censor bool) newrelic.Logger {
	return &Logger{
		l:     logger,
		debug: debug,
		censor: censor,
	}
}

// NewDefaultLogger returns a default Contextual Logger for birkirb's logger.
func NewDefaultLogger() newrelic.Logger {
	l := stdlib.NewDefaultLogger()
	return NewLogger(l, false, true)
}

// Error logs newrelic errors
func (l Logger) Error(msg string, fields map[string]interface{}) {
	l.l.WithFields(l.convert(fields)...).Error(msg)
}

// Warn logs newrelic warnings
func (l Logger) Warn(msg string, fields map[string]interface{}) {
	l.l.WithFields(l.convert(fields)...).Warn(msg)
}

// Info logs newrelic info messages
func (l Logger) Info(msg string, fields map[string]interface{}) {
	l.l.WithFields(l.convert(fields)...).Info(msg)
}

// Debug logs newrelic debug messages
func (l Logger) Debug(msg string, fields map[string]interface{}) {
	l.l.WithFields(l.convert(fields)...).Debug(msg)
}

// DebugEnabled returns whether debug is on
func (l Logger) DebugEnabled() bool {
	return l.debug
}
var filterRegEx = regexp.MustCompile(`[0-9a-z]+`)

func (l Logger) convert(c map[string]interface{}) []interface{} {
	output := make([]interface{}, 0, 2*len(c))
	for k, v := range c {
		 s := fmt.Sprint(v)
		 if k == "license_key" && l.censor {
			 s = filterRegEx.ReplaceAllString(s, "[REDACTED]")
		 }
		output = append(output, k, s)
	}
	return output
}
