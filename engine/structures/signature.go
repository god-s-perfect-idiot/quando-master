package structures

type CallSignature struct {
	MethodIdentifier string
	Parameters       []Parameter
	HasCallback      bool
}

type Parameter struct {
	Identifier string
	Type       string
	Value      interface{}
}

func NewCallSignature(methodIdentifier string, parameters []Parameter, hasCallback bool) *CallSignature {
	return &CallSignature{
		MethodIdentifier: methodIdentifier,
		Parameters:       parameters,
		HasCallback:      hasCallback,
	}
}

func (cs *CallSignature) GetMethodIdentifier() string {
	return cs.MethodIdentifier
}

func (cs *CallSignature) GetParameters() []Parameter {
	return cs.Parameters
}

func (cs *CallSignature) GetParameterCount() int {
	return len(cs.Parameters)
}

func (cs *CallSignature) GetParameter(index int) Parameter {
	return cs.Parameters[index]
}

func (cs *CallSignature) GetParameterIdentifier(index int) string {
	return cs.Parameters[index].Identifier
}

func (cs *CallSignature) GetParameterType(index int) string {
	return cs.Parameters[index].Type
}

func (cs *CallSignature) GetParameterValue(index int) interface{} {
	return cs.Parameters[index].Value
}

func (cs *CallSignature) CallbackExists() bool {
	return cs.HasCallback
}
