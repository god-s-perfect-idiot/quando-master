package structures

type CallNode struct {
	identifier string
	Definition Definition
	Method     Method
	Children   []*CallNode
	Parent     *CallNode
}

type CallGraph struct {
	Roots []*CallNode
}

func (c *CallGraph) appendChild(parent *CallNode, child *CallNode) {
	child.Parent = parent
	parent.Children = append(parent.Children, child)
}

func (c *CallGraph) GetRoots() []*CallNode {
	return c.Roots
}

// func (c *CallGraph) isTreeBald() bool {
// 	return c.Root == nil
// }

func (c *CallGraph) AttachMethodToNode(node *CallNode, methodIdentifier string, method Method) {
	if node.Definition.Signature.MethodIdentifier == methodIdentifier {
		if node.Method.Identifier == "" {
			node.Method = method
		}
	}
	for _, child := range node.Children {
		c.AttachMethodToNode(child, methodIdentifier, method)
	}
}

func (c *CallGraph) AttachMethod(methodIdentifier string, method Method) {
	for _, root := range c.Roots {
		c.AttachMethodToNode(root, methodIdentifier, method)
	}
}

func ConstructCallGraph(definitions []Definition) CallGraph {
	callStack := NewStack()
	callGraph := CallGraph{}
	for _, definition := range definitions {
		switch definition.Type {
		case "invocation":
			callnode := CallNode{
				identifier: definition.GetDefinitionString(),
				Definition: definition,
				Children:   []*CallNode{},
			}
			lastCall := callStack.Peek()
			if lastCall != nil {
				callGraph.appendChild(lastCall.(*CallNode), &callnode)
			} else {
				callGraph.Roots = append(callGraph.Roots, &callnode)
			}
			if definition.Signature.CallbackExists() {
				callStack.Push(&callnode)
			}
			// if callGraph.isTreeBald() {
			// 	callGraph.Root = &callnode
			// }
		case "callback terminator":
			callStack.Pop()
		default:
			continue
		}
	}
	return callGraph
}

func (c *CallGraph) Print(node *CallNode) {
	print(node.Definition.Signature.MethodIdentifier)
	for _, child := range node.Children {
		print(" -> ")
		c.Print(child)
	}
}

func (c *CallGraph) PrintAll() {
	print("Call Graph:\n")
	for _, root := range c.Roots {
		print("Root: ")
		c.Print(root)
		print("\n")
	}
}
