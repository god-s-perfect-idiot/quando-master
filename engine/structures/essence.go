package structures

import (
	"context"
)

type Essence struct {
	Context         context.Context
	CallStack       *Stack
	Invocations     *InvocationTable
	DependencyGraph *CallGraph
	Val             float64
	CallPipe        *chan map[string]interface{}
}

type CallData struct {
	Route string
	Body  []byte
}

func NewExecutionContext(invocations InvocationTable, dependencyGraph CallGraph) *Essence {
	scriptContext := context.Background()
	callStack := NewStack()
	return &Essence{
		Context:         scriptContext,
		CallStack:       callStack,
		Invocations:     &invocations,
		DependencyGraph: &dependencyGraph,
	}
}

func (e *Essence) ConnectCallPipe(channel *chan map[string]interface{}) {
	e.CallPipe = channel
}
