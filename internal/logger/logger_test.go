package logger

import "testing"

func TestLog(t *testing.T) {
	Debug("Call in TestLog", 123)
	Info("Call in TestLog", 321)
	Warn("Call in TestLog", 3.14)
	Error("Call in TestLog", true)
}
