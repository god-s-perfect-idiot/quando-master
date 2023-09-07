package analyser

import "testing"

func TestNewCursor(t *testing.T) {
	c := NewCursor()
	if c == nil {
		t.Error("NewCursor() should not return nil")
	}
}

func TestNewCursorControls(t *testing.T) {
	c := NewCursor()
	if c.line != 0 {
		t.Error("Cursor line should be 0")
	}
	if c.column != 0 {
		t.Error("Cursor column should be 0")
	}
}

func TestCursorAdvanceLine(t *testing.T) {
	c := NewCursor()
	c.advanceLine()
	if c.line != 1 {
		t.Error("Cursor line should be 1")
	}
}

func TestCursorAdvanceColumn(t *testing.T) {
	c := NewCursor()
	c.advanceColumn()
	if c.column != 1 {
		t.Error("Cursor column should be 1")
	}
}

func TestCursorGetL(t *testing.T) {
	c := NewCursor()
	if c.getL() != c.line {
		t.Error("Cursor line should be 0")
	}
}

func TestCursorGetC(t *testing.T) {
	c := NewCursor()
	if c.getC() != c.column {
		t.Error("Cursor column should be 0")
	}
}
