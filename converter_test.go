package newrelic

import (
	"testing"
)
func TestNoConversion(t *testing.T) {
	noChange:= "Post \"https://collector.newrelic.com/agent_listener/invoke_raw_method?license_key=1234567891234567891234567891231111111111&marshal_format=json&method=preconnect&protocol_version=17\": net/http: TLS handshake timeout"
	d := map[string]interface{}{
		"error": noChange,
	}
	l := Logger{
		censor: false,
	}
	output := l.convert(d)
	if output[1] != noChange {
		t.Errorf("No match. Found: `%s` Wanted: `%s`", output[1], noChange)
	}
}

func TestErrorConversion(t *testing.T) {
	d := map[string]interface{}{
		"error": "Post \"https://collector.newrelic.com/agent_listener/invoke_raw_method?license_key=1234567891234567891234567891231111111111&marshal_format=json&method=preconnect&protocol_version=17\": net/http: TLS handshake timeout",
	}
	l := Logger{
		censor: true,
	}
	output := l.convert(d)
	expected := "Post \"https://collector.newrelic.com/agent_listener/invoke_raw_method?license_key=[REDACTED]&marshal_format=json&method=preconnect&protocol_version=17\": net/http: TLS handshake timeout"
	if output[1] != expected {
		t.Errorf("No match. Found: `%s` Wanted: `%s`", output[1], expected)
	}
}
