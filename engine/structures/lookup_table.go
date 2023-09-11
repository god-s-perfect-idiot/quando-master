package structures

type LookupTable struct {
	Methods []Method
}

func NewLookupTable() *LookupTable {
	return &LookupTable{
		Methods: make([]Method, 0),
	}
}

func (l *LookupTable) Append(api Method) {
	l.Methods = append(l.Methods, api)
}

// TODO optimize this

func (l *LookupTable) GetAPI(identifier string) (Method, bool) {
	for _, method := range l.Methods {
		if method.Identifier == identifier {
			return method, true
		}
	}
	return Method{}, false
}

func ConstructLookupTable(methods []Method) *LookupTable {
	l := NewLookupTable()
	l.Methods = methods
	return l
}
