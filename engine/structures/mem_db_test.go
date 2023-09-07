package structures

import "testing"

func TestNewMemory(t *testing.T) {
	memory := NewMemory()
	if memory == nil {
		t.Error("Expected memory to be instantiated")
	}
}

func TestMemory_Get(t *testing.T) {
	memory := NewMemory()
	memory.Set("test", "value")
	if memory.Get("test") != "value" {
		t.Error("Expected memory to return value")
	}
}

func TestMemory_Set(t *testing.T) {
	memory := NewMemory()
	memory.Set("test", "value")
	if memory.Get("test") != "value" {
		t.Error("Expected memory to return value")
	}
}

func TestMemory_Delete(t *testing.T) {
	memory := NewMemory()
	memory.Set("test", "value")
	memory.Delete("test")
	if memory.Check("test") {
		t.Error("Expected memory to be deleted")
	}
}

func TestMemory_Clear(t *testing.T) {
	memory := NewMemory()
	memory.Set("test", "value")
	memory.Clear()
	if memory.Check("test") {
		t.Error("Expected memory to be cleared")
	}
}

func TestMemory_Check(t *testing.T) {
	memory := NewMemory()
	memory.Set("test", "value")
	if !memory.Check("test") {
		t.Error("Expected memory to check for key")
	}
}

func TestMemory_Check_False(t *testing.T) {
	memory := NewMemory()
	if memory.Check("test") {
		t.Error("Expected memory to not check for key")
	}
}

func TestMemory_Check_Nil(t *testing.T) {
	memory := NewMemory()
	if memory.Check("test") {
		t.Error("Expected memory to not check for key")
	}
}

func TestMemory_Check_Nil_False(t *testing.T) {
	memory := NewMemory()
	if memory.Check("test") {
		t.Error("Expected memory to not check for key")
	}
}
