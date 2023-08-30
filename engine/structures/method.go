package structures

import (
	"strconv"
)

type Method struct {
	Identifier string
	Function   interface{}
	//Only need to define those params that could have Val as a value
	Params   []Param
	Type     string
	Iterator bool
	Arbiter  bool
	DataKeys []string
}

type Param struct {
	Identifier string
	Type       string
}

func (m *Method) FindParam(id string) *Param {
	for _, param := range m.Params {
		if param.Identifier == id {
			return &param
		}
	}
	return nil
}

func (m *Method) cleanParam(param Parameter, essence *Executable) Parameter {
	switch param.Type {
	case "STRING":
		param.Value = param.Value.(string)[1 : len(param.Value.(string))-1]
	case "BOOLEAN":
		param.Value = param.Value == "true"
	case "NUMBER":
		param.Value, _ = strconv.Atoi(param.Value.(string))
	case "FLOAT":
		param.Value, _ = strconv.ParseFloat(param.Value.(string), 64)
	case "VAL":
		param.Value = essence.Val
		parameter := m.FindParam(param.Identifier)
		if parameter != nil {
			switch parameter.Type {
			case "STRING":
				param.Value = strconv.FormatFloat(param.Value.(float64), 'f', -1, 64)
			}
		}
	}
	return param
}

func (m *Method) Call(parameters []Parameter, executable *Executable) {
	params := make(map[string]interface{})
	params["callPipe"] = executable.CallPipe
	params["val"] = executable.Val
	for _, parameter := range parameters {
		// Only binding early now. Fetch val from context later
		parameter = m.cleanParam(parameter, executable)
		params[parameter.Identifier] = parameter.Value
	}
	if m.IsArbiter() {
		params["keys"] = executable.Data["keys"].([]int)
		params["nodeCount"] = executable.Data["nodeCount"].(int)
		params["sequence"] = executable.Data["sequence"].([]int)
	}
	if m.DataKeys != nil {
		for _, key := range m.DataKeys {
			if executable.Data[key] != nil {
				params[key] = executable.Data[key]
			}
		}
	}
	val, data := m.Function.(func(map[string]interface{}) (float64, map[string]interface{}))(params)
	if m.Type == "callback" && val != -1.0 {
		executable.Val = val
	}
	if data != nil {
		for k, v := range data {
			executable.SetData(k, v)
		}
	}
}

func (m *Method) IsIterator() bool {
	return m.Iterator
}

func (m *Method) IsArbiter() bool {
	return m.Arbiter
}
