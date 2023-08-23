package structures

import "testing"

func TestNewLineBreak(t *testing.T) {
	l := NewLineBreak(0)
	if l == nil {
		t.Error("NewLineBreak() should not return nil")
	}
}

func TestNewLineBreakControls(t *testing.T) {
	l := NewLineBreak(0)
	if l.Line != 0 {
		t.Error("LineBreak line should be 0")
	}
	if l.Type != "line break" {
		t.Error("LineBreak type should be line break")
	}
}

func TestNewCallbackterminator(t *testing.T) {
	c := NewCallbackterminator(0)
	if c == nil {
		t.Error("NewCallbackterminator() should not return nil")
	}
}

func TestNewCallbackterminatorControls(t *testing.T) {
	c := NewCallbackterminator(0)
	if c.Line != 0 {
		t.Error("Callbackterminator line should be 0")
	}
	if c.Type != "callback terminator" {
		t.Error("Callbackterminator type should be callback terminator")
	}
}

func TestNewInvocation(t *testing.T) {
	i := NewInvocation(CallSignature{
		MethodIdentifier: "test",
		Parameters:       make([]Parameter, 0),
		HasCallback:      false,
	}, 0)
	if i == nil {
		t.Error("NewInvocation() should not return nil")
	}
}

func TestNewInvocationControls(t *testing.T) {
	i := NewInvocation(CallSignature{
		MethodIdentifier: "test",
		Parameters:       make([]Parameter, 0),
		HasCallback:      false,
	}, 0)
	if i.Line != 0 {
		t.Error("Invocation line should be 0")
	}
	if i.Type != "invocation" {
		t.Error("Invocation type should be invocation")
	}
	if i.Signature.MethodIdentifier != "test" {
		t.Error("Invocation signature method identifier should be test")
	}
}

func TestNewInvocationWithCallbackControls(t *testing.T) {
	i := NewInvocation(CallSignature{
		MethodIdentifier: "test",
		Parameters:       make([]Parameter, 0),
		HasCallback:      true,
	}, 0)
	if i.Line != 0 {
		t.Error("Invocation line should be 0")
	}
	if i.Type != "invocation" {
		t.Error("Invocation type should be invocation")
	}
	if i.Signature.MethodIdentifier != "test" {
		t.Error("Invocation signature method identifier should be test")
	}
	if !i.Signature.HasCallback {
		t.Error("Invocation signature should have callback")
	}
}

func TestGetDefinitionString(t *testing.T) {
	i := NewInvocation(CallSignature{
		MethodIdentifier: "test",
		Parameters:       make([]Parameter, 0),
		HasCallback:      true,
	}, 0)
	if i.GetDefinitionString() != "0:callback:test" {
		t.Error("Invocation definition string should be 0:callback:test")
	}
}

func TestGetDefinitionStringWithoutCallback(t *testing.T) {
	i := NewInvocation(CallSignature{
		MethodIdentifier: "test",
		Parameters:       make([]Parameter, 0),
		HasCallback:      false,
	}, 0)
	if i.GetDefinitionString() != "0:action:test" {
		t.Error("Invocation definition string should be 0:action:test")
	}
}
