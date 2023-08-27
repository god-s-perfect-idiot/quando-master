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
	Data            map[string]interface{}
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
		Data:            make(map[string]interface{}),
	}
}

func (e *Essence) ConnectCallPipe(channel *chan map[string]interface{}) {
	e.CallPipe = channel
}

func (e *Essence) SetData(key string, value interface{}) {
	e.Data[key] = value
}

func (e *Essence) GetData(key string) interface{} {
	return e.Data[key]
}

func (e *Essence) ValidateData(key string) bool {
	_, ok := e.Data[key]
	return ok
}
