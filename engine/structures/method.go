package structures

import (
	"strconv"
)

type Method struct {
	Identifier string
	Function   interface{}
	Type       string
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
		if parameter.Value == "val" {
			params[parameter.Identifier] = 0.5
		} else {
			params[parameter.Identifier] = parameter.Value
		}
	}
	val := m.Function.(func(map[string]interface{}) float64)(params)
	if m.Type == "callback" {
		essence.Val = val
	}
}
