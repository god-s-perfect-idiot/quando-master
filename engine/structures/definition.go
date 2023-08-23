package structures

import (
	"fmt"
)

type Definition struct {
	Type      string
	Signature CallSignature
	Line      int
}

func NewLineBreak(lineCount int) *Definition {
	return &Definition{
		Type: "line break",
		Line: lineCount,
	}
}

func NewCallbackterminator(lineCount int) *Definition {
	return &Definition{
		Type: "callback terminator",
		Line: lineCount,
	}
}

func NewInvocation(signature CallSignature, lineCount int) *Definition {
	return &Definition{
		Type:      "invocation",
		Signature: signature,
		Line:      lineCount,
	}
}

func (d *Definition) GetDefinitionString() string {
	if d.Signature.HasCallback {
		return fmt.Sprintf("%d:callback:%s", d.Line, d.Signature.GetMethodIdentifier())
	}
	return fmt.Sprintf("%d:action:%s", d.Line, d.Signature.GetMethodIdentifier())
}
