package structures

type InvocationTable []Invocation

type Invocation struct {
	Identifier string
	Signature  CallSignature
	Type       string
}

func newInvocationTable() *InvocationTable {
	return &InvocationTable{}
}

func (i *InvocationTable) append(invocation Invocation) {
	*i = append(*i, invocation)
}

func GenerateInvocationTable(definitions []Definition) InvocationTable {
	invocationTable := newInvocationTable()
	for _, definition := range definitions {
		if definition.Type == "invocation" {
			if definition.Signature.CallbackExists() {
				invocationTable.append(Invocation{
					Identifier: definition.GetDefinitionString(),
					Signature:  definition.Signature,
					Type:       "callback",
				})
			} else {
				invocationTable.append(Invocation{
					Identifier: definition.GetDefinitionString(),
					Signature:  definition.Signature,
					Type:       "action",
				})
			}
		}
	}
	return *invocationTable
}
