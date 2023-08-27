package structures

import (
	"strconv"
)

type Method struct {
	Identifier string
	Function   interface{}
	Type       string
	Iterator   bool
	Arbiter    bool
}

func (m *Method) cleanParam(param Parameter, essence *Essence) Parameter {
	switch param.Type {
	case "STRING":
		param.Value = param.Value.(string)[1 : len(param.Value.(string))-1]
	case "BOOLEAN":
		param.Value = param.Value == "true"
	case "NUMBER":
		param.Value, _ = strconv.Atoi(param.Value.(string))
	case "VAL":
		param.Value = essence.Val
	}
	return param
}

func (m *Method) Call(parameters []Parameter, essence *Essence) {
	params := make(map[string]interface{})
	params["callPipe"] = essence.CallPipe
	for _, parameter := range parameters {
		// Only binding early now. Fetch val from context later
		parameter = m.cleanParam(parameter, essence)
		params[parameter.Identifier] = parameter.Value
	}
	if m.IsArbiter() {
		params["keys"] = essence.Data["keys"].([]int)
		params["nodeCount"] = essence.Data["nodeCount"].(int)
		params["sequence"] = essence.Data["sequence"].([]int)
	}
	val, data := m.Function.(func(map[string]interface{}) (float64, map[string]interface{}))(params)
	if m.Type == "callback" {
		essence.Val = val
	}
	if data != nil {
		for k, v := range data {
			essence.SetData(k, v)
		}
	}
}

func (m *Method) IsIterator() bool {
	return m.Iterator
}

func (m *Method) IsArbiter() bool {
	return m.Arbiter
}
