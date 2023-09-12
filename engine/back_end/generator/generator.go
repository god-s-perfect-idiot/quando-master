package generator

import (
	"quando/engine/structures"
)

type Generator struct {
	Lookups Lookups
}

func NewGenerator() *Generator {
	return &Generator{
		Lookups: *GenerateLookups(),
	}
}

func (g *Generator) AttachHandler(essence structures.Executable) {
	for _, invocation := range *essence.Invocations {
		switch invocation.Type {
		case "callback":
			api, valid := g.Lookups.Callbacks.GetAPI(invocation.Signature.MethodIdentifier)
			if valid {
				essence.DependencyGraph.AttachMethod(api.Identifier, api)
			}
		case "action":
			api, valid := g.Lookups.Actions.GetAPI(invocation.Signature.MethodIdentifier)
			if valid {
				essence.DependencyGraph.AttachMethod(api.Identifier, api)
			}
		case "conditional callback":
			api, valid := g.Lookups.Callbacks.GetAPI(invocation.Signature.MethodIdentifier)
			if valid {
				essence.DependencyGraph.AttachMethod(api.Identifier, api)
			}
		default:
			continue
		}
	}
}
