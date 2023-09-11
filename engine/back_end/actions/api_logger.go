package actions

import (
	"quando/engine/structures"
)

type LoggerClient struct{}

func GetLoggerActions() []structures.Method {
	loggerClient := NewLogger()
	return []structures.Method{
		{
			Identifier: "quando.log",
			Function:   loggerClient.Log,
			Params: []structures.Param{
				{
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

func (l *LoggerClient) Log(params map[string]interface{}, _ *structures.RunContext) {
	message := params["text"].(string)
	l.log(message)
}

func (l *LoggerClient) log(message string) {
	println(message)
}
