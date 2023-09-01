package actions

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger()
	if l == nil {
		t.Error("NewLogger() should not return nil")
	}
}

func TestLoggerGetLoggerActions(t *testing.T) {
	loggerActions := GetLoggerActions()
	if len(loggerActions) == 0 {
		t.Error("No logger actions found")
	}
}

func TestLogAction(t *testing.T) {
	loggerActions := GetLoggerActions()
	log := loggerActions[0]
	if log.Identifier != "quando.log" {
		t.Error("Identifier should be quando.log")
	}
	if log.Type != "action" {
		t.Error("Type should be action")
	}
}
