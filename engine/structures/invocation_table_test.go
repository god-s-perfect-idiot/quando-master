package structures

import "testing"

func TestNewInvocationTable(t *testing.T) {
	i := newInvocationTable()
	if i == nil {
		t.Error("NewInvocationTable() should not return nil")
	}
}

func TestGenerateInvocationTable(t *testing.T) {
	d := []Definition{
		Definition{
			Line: 0,
			Signature: CallSignature{
				MethodIdentifier: "test",
				Parameters:       make([]Parameter, 0),
				HasCallback:      false,
			},
			Type: "invocation",
		},
	}
	i := GenerateInvocationTable(d)
	if i == nil {
		t.Error("GenerateInvocationTable() should not return nil")
	}
	if len(i) != 1 {
		t.Error("GenerateInvocationTable() should return 1 invocation")
	}
	if i[0].Identifier != "0:action:test" {
		t.Error("GenerateInvocationTable() should return invocation with identifier test")
	}
	if i[0].Signature.MethodIdentifier != "test" {
		t.Error("GenerateInvocationTable() should return invocation with signature method identifier test")
	}
	if i[0].Signature.HasCallback {
		t.Error("GenerateInvocationTable() should return invocation with signature has callback false")
	}
	if i[0].Type != "action" {
		t.Error("GenerateInvocationTable() should return invocation with type action")
	}
}

func TestGenerateInvocationTableWithCallback(t *testing.T) {
	d := []Definition{
		Definition{
			Line: 0,
			Signature: CallSignature{
				MethodIdentifier: "test",
				Parameters:       make([]Parameter, 0),
				HasCallback:      true,
			},
			Type: "invocation",
		},
	}
	i := GenerateInvocationTable(d)
	if i == nil {
		t.Error("GenerateInvocationTable() should not return nil")
	}
	if len(i) != 1 {
		t.Error("GenerateInvocationTable() should return 1 invocation")
	}
	if i[0].Identifier != "0:callback:test" {
		t.Error("GenerateInvocationTable() should return invocation with identifier test")
	}
	if i[0].Signature.MethodIdentifier != "test" {
		t.Error("GenerateInvocationTable() should return invocation with signature method identifier test")
	}
	if !i[0].Signature.HasCallback {
		t.Error("GenerateInvocationTable() should return invocation with signature has callback false")
	}
}

func TestGenerateInvocationTableWithMultipleInvocations(t *testing.T) {
	d := []Definition{
		Definition{
			Line: 0,
			Signature: CallSignature{
				MethodIdentifier: "test",
				Parameters:       make([]Parameter, 0),
				HasCallback:      false,
			},
			Type: "invocation",
		},
		Definition{
			Line: 1,
			Signature: CallSignature{
				MethodIdentifier: "test",
				Parameters:       make([]Parameter, 0),
				HasCallback:      false,
			},
			Type: "invocation",
		},
	}
	i := GenerateInvocationTable(d)
	if i == nil {
		t.Error("GenerateInvocationTable() should not return nil")
	}
	if len(i) != 2 {
		t.Error("GenerateInvocationTable() should return 2 invocations")
	}
}

func TestAppendInvocation(t *testing.T) {
	i := newInvocationTable()
	i.append(Invocation{
		Identifier: "0:action:test",
		Signature: CallSignature{
			MethodIdentifier: "test",
			Parameters:       make([]Parameter, 0),
			HasCallback:      false,
		},
		Type: "action",
	})
	if len(*i) != 1 {
		t.Error("appendInvocation() should append invocation")
	}
}
