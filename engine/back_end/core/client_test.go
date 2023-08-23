package core

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient()
	if c == nil {
		t.Error("NewClient() should not return nil")
	}
}
