package structures

type CallNode struct {
	identifier   string
	Definition   Definition
	Type         string
	Method       Method
	MainChildren []*CallNode
	AltChildren  []*CallNode
	Parent       *CallNode
}

type CallGraph struct {
	Roots []*CallNode
}

func (c *CallGraph) appendChild(parent *CallNode, child *CallNode, isAlt bool) {
	child.Parent = parent
	if parent.Type == "conditional callback" {
		if isAlt {
			parent.AltChildren = append(parent.AltChildren, child)
		} else {
			parent.MainChildren = append(parent.MainChildren, child)
		}
	} else {
		parent.MainChildren = append(parent.MainChildren, child)
	}
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
	for _, child := range node.MainChildren {
		c.AttachMethodToNode(child, methodIdentifier, method)
	}
	for _, child := range node.AltChildren {
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
	alt := false
	for _, definition := range definitions {
		switch definition.Type {
		case "invocation":
			callnode := CallNode{
				identifier:   definition.GetDefinitionString(),
				Definition:   definition,
				MainChildren: []*CallNode{},
			}
			lastCall := callStack.Peek()
			if lastCall != nil {
				callGraph.appendChild(lastCall.(*CallNode), &callnode, alt)
			} else {
				callGraph.Roots = append(callGraph.Roots, &callnode)
			}
			switch definition.Type {
			case "invocation":
				if definition.Signature.HasCallback {
					callnode.Type = "callback"
					callStack.Push(&callnode)
				} else {
					callnode.Type = "action"
				}
			}
			// if callGraph.isTreeBald() {
			// 	callGraph.Root = &callnode
			// }
		case "callback terminator":
			node := callStack.Pop().(*CallNode)
			if node.Type == "conditional callback" {
				alt = false
			}
		case "conditional callback":
			alt = true
			parentNode := callStack.Peek().(*CallNode)
			parentNode.Type = "conditional callback"
		default:
			continue
		}
	}
	return callGraph
}

func (c *CallGraph) Print(node *CallNode) {
	print(node.Definition.Signature.MethodIdentifier)
	for _, child := range node.MainChildren {
		print(" -> ")
		c.Print(child)
	}
	for _, child := range node.AltChildren {
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
