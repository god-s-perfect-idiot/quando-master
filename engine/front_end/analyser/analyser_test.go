package analyser

import "testing"

var sourceCode = `
	quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
		quando.control.key ch="a", upDown="down", onOff=val
	}
`

func TestNewAnalyser(t *testing.T) {
	a := NewAnalyser(sourceCode)
	if a == nil {
		t.Error("NewAnalyser() should not return nil")
	}
}

func TestNewAnalyserCursor(t *testing.T) {
	a := NewAnalyser(sourceCode)
	if &a.cursor == nil {
		t.Error("NewAnalyser() should not return nil")
	}
}

func TestNewAnalyserScanner(t *testing.T) {
	a := NewAnalyser(sourceCode)
	if &a.scanner == nil {
		t.Error("NewAnalyser() should not return nil")
	}
}

func TestScan(t *testing.T) {
	a := NewAnalyser(sourceCode)
	essence := a.Scan()
	if essence == nil {
		t.Error("Scan() should not return nil")
	}
	if essence.Invocations == nil {
		t.Error("Scan() should not return nil invocations for the source code")
	}
	if len(*essence.Invocations) == 0 {
		t.Error("Scan() should not return empty invocations for the source code")
	}
	if essence.DependencyGraph == nil {
		t.Error("Scan() should not return nil dependency graph for the source code")
	}
	if (*essence.DependencyGraph).Roots == nil {
		t.Error("Scan() should not return nil roots for the source code")
	}
	if len((*essence.DependencyGraph).Roots) != 1 {
		t.Error("Scan() should return one root for the source code")
	}
}

func TestTokenize(t *testing.T) {
	a := NewAnalyser(sourceCode)
	definitionLB := a.tokenize()
	if definitionLB == nil {
		t.Error("tokenize() should not return nil")
	}
	if definitionLB.Line != 0 {
		t.Error("tokenize() should return line 0")
	}
	if definitionLB.Type != "line break" {
		t.Error("tokenize() should return line break type")
	}
}

func TestTokenizeInvocation(t *testing.T) {
	a := NewAnalyser(sourceCode)
	a.scanner.line = `quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {`
	definitionInv := a.tokenize()
	if definitionInv == nil {
		t.Error("tokenize() should not return nil")
	}
	if definitionInv.Line != 0 {
		t.Error("tokenize() should return line 0")
	}
	if definitionInv.Type != "invocation" {
		t.Error("tokenize() should return invocation type")
	}
	if definitionInv.Signature.MethodIdentifier != "quando.key.handleKey" {
		t.Error("tokenize() should return method identifier quando.key.handleKey")
	}
}
