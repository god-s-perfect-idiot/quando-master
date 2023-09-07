package generator

import "testing"

func TestGenerateLookups(t *testing.T) {
	lookups := GenerateLookups()
	if lookups == nil {
		t.Error("GenerateLookups() should not return nil")
	}
	if lookups.Callbacks == nil {
		t.Error("Callbacks should not return nil")
	}
	if lookups.Actions == nil {
		t.Error("Actions should not return nil")
	}
	if lookups.Callbacks.Methods == nil {
		t.Error("Callbacks.Methods should not return nil")
	}
	if lookups.Actions.Methods == nil {
		t.Error("Actions.Methods should not return nil")
	}
	if len(lookups.Callbacks.Methods) == 0 {
		t.Error("Callbacks.Methods should not be empty")
	}
	if len(lookups.Actions.Methods) == 0 {
		t.Error("Actions.Methods should not be empty")
	}
}
