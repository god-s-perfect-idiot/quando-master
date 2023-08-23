package analyser

import "testing"

var scannerCode = `
	line 1
	line 2
	line 3
`

var scannerLine = "line 1"

var scannerSplitLine = "split here, and=here"

func TestNewScanner(t *testing.T) {
	s := NewScanner(scannerCode)
	if s == nil {
		t.Error("NewScanner() should not return nil")
	}
}

func TestNewScannerControls(t *testing.T) {
	s := NewScanner(scannerCode)
	if s.scanner == nil {
		t.Error("Scanner should not return nil")
	}
	if s.line != "" {
		t.Error("Scanner line should be 0")
	}
}

func TestSetLine(t *testing.T) {
	s := NewScanner(scannerCode)
	s.setLine("line 1")
	if s.line != "line 1" {
		t.Error("Scanner line should be 'line 1'")
	}
}

func TestGetLineLength(t *testing.T) {
	s := NewScanner(scannerCode)
	s.setLine("line 1")
	if s.getLineLength() != 6 {
		t.Error("Scanner line length should be 6")
	}
}

func TestGetLine(t *testing.T) {
	s := NewScanner(scannerCode)
	s.setLine("line 1")
	if s.getLine() != s.line {
		t.Error("Scanner line should be 'line 1'")
	}
}

func TestReadLine(t *testing.T) {
	s := NewScanner(scannerLine)
	for s.Scan() {
		valid := s.readLine()
		if !valid {
			t.Error("Scanner should read line")
		}
	}
	if s.line != "line 1" {
		t.Error("Scanner line should be 'line 1'")
	}
}

func TestSplit(t *testing.T) {
	s := NewScanner(scannerSplitLine)
	for s.Scan() {
		valid := s.readLine()
		if !valid {
			t.Error("Scanner should read line")
		}
	}
	parts := s.split()
	if len(parts) != 4 {
		t.Error("Scanner should split line into 4 parts")
	}
	if parts[0] != "split" {
		t.Error("Scanner should split line into 'split'")
	}
	if parts[1] != "here" {
		t.Error("Scanner should split line into 'here'")
	}
	if parts[2] != "and" {
		t.Error("Scanner should split line into 'and'")
	}
	if parts[3] != "here" {
		t.Error("Scanner should split line into 'here'")
	}
}
