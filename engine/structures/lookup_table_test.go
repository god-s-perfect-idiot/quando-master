package structures

import "testing"

func TestNewLookupTable(t *testing.T) {
	l := NewLookupTable()
	if l == nil {
		t.Error("NewLookupTable() should not return nil")
	}
}

func TestNewLookupTableControls(t *testing.T) {
	l := NewLookupTable()
	if l.Methods == nil {
		t.Error("LookupTable table should not be nil")
	}
}

func TestLookupTableAppend(t *testing.T) {
	l := NewLookupTable()
	l.Append(Method{})
	if len(l.Methods) != 1 {
		t.Error("LookupTable should have one method")
	}
}

func TestLookupTableGetAPI(t *testing.T) {
	l := NewLookupTable()
	l.Append(Method{
		Identifier: "test",
	})
	method, ok := l.GetAPI("test")
	if !ok {
		t.Error("LookupTable should return ok")
	}
	if method.Identifier != "test" {
		t.Error("LookupTable should return method with identifier test")
	}
}

func TestConstructLookupTable(t *testing.T) {
	l := ConstructLookupTable([]Method{
		Method{
			Identifier: "test",
		},
	})
	if l == nil {
		t.Error("ConstructLookupTable() should not return nil")
	}
	if len(l.Methods) != 1 {
		t.Error("ConstructLookupTable() should return one method")
	}
	if l.Methods[0].Identifier != "test" {
		t.Error("ConstructLookupTable() should return method with identifier test")
	}
}
