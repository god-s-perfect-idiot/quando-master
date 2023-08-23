package structures

import "testing"

func TestNewStack(t *testing.T) {
	s := NewStack()
	if s == nil {
		t.Error("NewStack() should not return nil")
	}
}

func TestNewStackControls(t *testing.T) {
	s := NewStack()
	if s.elements == nil {
		t.Error("Stack elements should not be nil")
	}
	if s.pointer != -1 {
		t.Error("Stack pointer should be -1")
	}
}

func TestStackPush(t *testing.T) {
	s := NewStack()
	s.Push(1)
	if len(s.elements) != 1 {
		t.Error("Stack should have one element")
	}
	if s.pointer != 0 {
		t.Error("Stack pointer should be 0")
	}
}

func TestStackPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	element := s.Pop()
	if element != 1 {
		t.Error("Stack should return 1")
	}
	if len(s.elements) != 0 {
		t.Error("Stack should have zero elements")
	}
	if s.pointer != -1 {
		t.Error("Stack pointer should be -1")
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := NewStack()
	element := s.Pop()
	if element != nil {
		t.Error("Stack should return nil")
	}
	if len(s.elements) != 0 {
		t.Error("Stack should have zero elements")
	}
	if s.pointer != -1 {
		t.Error("Stack pointer should be -1")
	}
}

func TestStackPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	element := s.Peek()
	if element != 1 {
		t.Error("Stack should return 1")
	}
	if len(s.elements) != 1 {
		t.Error("Stack should have one element")
	}
	if s.pointer != 0 {
		t.Error("Stack pointer should be 0")
	}
}

func TestStackPeekEmpty(t *testing.T) {
	s := NewStack()
	element := s.Peek()
	if element != nil {
		t.Error("Stack should return nil")
	}
	if len(s.elements) != 0 {
		t.Error("Stack should have zero elements")
	}
	if s.pointer != -1 {
		t.Error("Stack pointer should be -1")
	}
}

func TestStackPeekEmptyAfterPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Pop()
	element := s.Peek()
	if element != nil {
		t.Error("Stack should return nil")
	}
	if len(s.elements) != 0 {
		t.Error("Stack should have zero elements")
	}
	if s.pointer != -1 {
		t.Error("Stack pointer should be -1")
	}
}

func TestStackPeekEmptyAfterPushPop(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Pop()
	s.Push(1)
	element := s.Peek()
	if element != 1 {
		t.Error("Stack should return 1")
	}
	if len(s.elements) != 1 {
		t.Error("Stack should have one element")
	}
	if s.pointer != 0 {
		t.Error("Stack pointer should be 0")
	}
}
