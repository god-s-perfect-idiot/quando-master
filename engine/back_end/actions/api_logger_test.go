package actions

import "testing"

func TestNewLogger(t *testing.T) {
	l := NewLogger()
	if l == nil {
		t.Error("NewLogger() should not return nil")
	}
}

func TestLoggerClient_Log(t *testing.T) {
	l := NewLogger()
	params := map[string]interface{}{
		"text": "test",
	}
	val, _ := l.Log(params)
	if val != 0.0 {
		t.Error("Log() should return 0.0")
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
