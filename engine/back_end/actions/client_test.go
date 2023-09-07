package actions

import (
	"quando/engine/structures"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Error("NewClient() should not return nil")
	}
}

func TestClientMethods(t *testing.T) {
	c := NewClient()
	if c.Methods == nil {
		t.Error("Client methods should not be nil")
	}
	if len(c.Methods) == 0 {
		t.Error("Client methods should not be empty")
	}
}

func TestClientGetActions(t *testing.T) {
	c := NewClient()
	actions := c.GetActions()
	if actions == nil {
		t.Error("Client actions should not be nil")
	}
	if len(actions) == 0 {
		t.Error("Client actions should not be empty")
	}
	if len(actions) != len(c.Methods) {
		t.Error("Client actions should be the same as client methods")
	}
}

func TestClientAddActions(t *testing.T) {
	c := NewClient()
	c.Methods = make([]structures.Method, 0)
	method := structures.Method{
		Identifier: "quando.test",
		Function: func(params map[string]interface{}) float64 {
			return 0.0
		},
		Type: "action",
	}
	c.AddActions([]structures.Method{method})
	if len(c.Methods) != 1 {
		t.Error("Client methods should have one method")
	}
	if c.Methods[0].Identifier != "quando.test" {
		t.Error("Client method should have identifier quando.test")
	}
}
