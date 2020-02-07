package newrelic

import (
	"bytes"
	syslog "log"
	"regexp"
	"testing"

	newrelic "github.com/newrelic/go-agent"
	"gopkg.in/birkirb/loggers.v1/mappers/stdlib"
)

func TestLoggersInterface(t *testing.T) {
	var _ newrelic.Logger = NewDefaultLogger()
}

func TestLoggersLevelOutput(t *testing.T) {
	l, b := newBufferedNRLog()
	l.Info("This is a test", nil)

	expectedMatch := "(?i)info.*This is a test"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}

func TestNRWithFieldsOutput(t *testing.T) {
	l, b := newBufferedNRLog()
	d := map[string]interface{}{
		"test": true,
	}
	l.Warn("This is a message", d)

	expectedMatch := "(?i)warn.*This is a message.*test.*=true"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}


func TestNRWithRedacting(t *testing.T) {
	l, b := newBufferedNRLog()
	d := map[string]interface{}{
		"license_key": "123456abcd",
	}
	l.Warn("This is a message", d)

	expectedMatch := `(?i)warn.*This is a message.*license_key=\[REDACTED\].*`
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}


func TestNRWithNoRedacting(t *testing.T) {
	l, b := newBufferedNRLogUncensored()
	d := map[string]interface{}{
		"license_key": "123456abcd",
	}
	l.Warn("This is a message", d)

	expectedMatch := `(?i)warn.*This is a message.*license_key=123456abcd.*`
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch)
	}
}


func TestNRWithMultipleFieldsOutput(t *testing.T) {
	l, b := newBufferedNRLog()
	d := map[string]interface{}{
		"test":  true,
		"Error": "serious",
	}
	l.Error("This is a message", d)

	// maps are random by design, so we can't know what order the fields will be in
	expectedMatch1 := "(?i)erro.*This is a message"
	expectedMatch2 := "test.*=true"
	expectedMatch3 := "Error.*=serious"
	actual := b.String()
	if ok, _ := regexp.Match(expectedMatch1, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch1)
	}
	if ok, _ := regexp.Match(expectedMatch2, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch2)
	}
	if ok, _ := regexp.Match(expectedMatch3, []byte(actual)); !ok {
		t.Errorf("Log output mismatch %s (actual) != %s (expected)", actual, expectedMatch3)
	}
}

func newBufferedNRLog() (newrelic.Logger, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)
	sl := syslog.New(bb, "", 0)
	l := stdlib.NewLogger(sl)
	return NewLogger(l, true, true), bb
}

func newBufferedNRLogUncensored() (newrelic.Logger, *bytes.Buffer) {
	var b []byte
	var bb = bytes.NewBuffer(b)
	sl := syslog.New(bb, "", 0)
	l := stdlib.NewLogger(sl)
	return NewLogger(l, true, false), bb
}