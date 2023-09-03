package structures

type Memory struct {
	Records map[string]interface{}
}

func NewMemory() *Memory {
	return &Memory{
		Records: make(map[string]interface{}),
	}
}

func (m *Memory) Get(key string) interface{} {
	return m.Records[key]
}

func (m *Memory) Set(key string, value interface{}) {
	m.Records[key] = value
}

func (m *Memory) Delete(key string) {
	delete(m.Records, key)
}

func (m *Memory) Clear() {
	m.Records = make(map[string]interface{})
}

func (m *Memory) Check(key string) bool {
	_, ok := m.Records[key]
	return ok
}
