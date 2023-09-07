package devices

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

func TestClientCallbacks(t *testing.T) {
	c := NewClient()
	if c.Methods == nil {
		t.Error("Client callbacks should not be nil")
	}
	if len(c.Methods) == 0 {
		t.Error("Client callbacks should not be empty")
	}
}

func TestClientGetCallbacks(t *testing.T) {
	c := NewClient()
	callbacks := c.GetCallbacks()
	if callbacks == nil {
		t.Error("Client callbacks should not be nil")
	}
	if len(callbacks) == 0 {
		t.Error("Client callbacks should not be empty")
	}
	if len(callbacks) != len(c.Methods) {
		t.Error("Client callbacks should be the same as client methods")
	}
}

func TestClientAddCallbacks(t *testing.T) {
	c := NewClient()
	c.Methods = make([]structures.Method, 0)
	method := structures.Method{
		Identifier: "quando.test",
		Function: func(params map[string]interface{}) float64 {
			return 0.0
		},
		Type: "callback",
	}
	c.AddCallbacks([]structures.Method{method})
	if len(c.Methods) != 1 {
		t.Error("Client methods should have one method")
	}
	if c.Methods[0].Identifier != "quando.test" {
		t.Error("Client method should have identifier quando.test")
	}
}
