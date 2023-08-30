package actions

import (
	"quando/engine/structures"
)

type LoggerClient struct{}

func GetLoggerActions() []structures.Method {
	loggerClient := NewLogger()
	return []structures.Method{
		structures.Method{
			Identifier: "quando.log",
			Function:   loggerClient.Log,
			Params: []structures.Param{
				structures.Param{
					Identifier: "text",
					Type:       "STRING",
				},
			},
			Type:     "action",
			Iterator: false,
			Arbiter:  false,
		},
	}
}

func NewLogger() *LoggerClient {
	return &LoggerClient{}
}

func (l *LoggerClient) Log(params map[string]interface{}) (float64, map[string]interface{}) {
	message := params["text"].(string)
	l.log(message)
	return 0.0, nil
}

func (l *LoggerClient) log(message string) {
	println(message)
}
