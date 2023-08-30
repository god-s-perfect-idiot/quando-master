package structures

import (
	"context"
)

type Executable struct {
	Hash            string
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

func NewExecutionContext(hash string, invocations InvocationTable, dependencyGraph CallGraph) *Executable {
	scriptContext := context.Background()
	callStack := NewStack()
	return &Executable{
		Hash:            hash,
		Context:         scriptContext,
		CallStack:       callStack,
		Invocations:     &invocations,
		DependencyGraph: &dependencyGraph,
		Data:            make(map[string]interface{}),
		Val:             0.5,
	}
}

func (e *Executable) ConnectCallPipe(channel *chan map[string]interface{}) {
	e.CallPipe = channel
}

func (e *Executable) SetData(key string, value interface{}) {
	e.Data[key] = value
}

func (e *Executable) GetData(key string) interface{} {
	return e.Data[key]
}

func (e *Executable) ValidateData(key string) bool {
	_, ok := e.Data[key]
	return ok
}
