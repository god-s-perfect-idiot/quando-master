package structures

type Stack struct {
	elements []interface{}
	pointer  int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		pointer:  -1,
	}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
	s.pointer++
}

func (s *Stack) Pop() interface{} {
	if s.pointer == -1 {
		return nil
	}
	element := s.elements[s.pointer]
	s.elements = s.elements[:s.pointer]
	s.pointer--
	return element
}

func (s *Stack) Peek() interface{} {
	if s.pointer == -1 {
		return nil
	}
	return s.elements[s.pointer]
}
