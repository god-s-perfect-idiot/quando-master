package structures

import "testing"

func TestNewExecutionContext(t *testing.T) {
	i := InvocationTable{}
	g := CallGraph{}
	e := NewExecutionContext("", i, g)
	if e == nil {
		t.Error("NewExecutionContext() should not return nil")
	}
}

func TestNewExecutionContextControls(t *testing.T) {
	i := InvocationTable{}
	g := CallGraph{}
	e := NewExecutionContext("", i, g)
	if e.Context == nil {
		t.Error("Context should not be nil")
	}
	if e.CallStack == nil {
		t.Error("CallStack should not be nil")
	}
	if e.Invocations == nil {
		t.Error("Invocations should not be nil")
	}
	if e.DependencyGraph == nil {
		t.Error("DependencyGraph should not be nil")
	}
	if e.Val != 0.5 {
		t.Error("Val should be 0.5")
	}
}
