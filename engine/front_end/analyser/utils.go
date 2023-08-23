package analyser

import (
	"strconv"
)

func IsBooleanParameter(value string) bool {
	return value == "true" || value == "false"
}

func IsIntegerParameter(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func IsFloatParameter(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func IsStringParameter(value string) bool {
	return value[0] == '"' && value[len(value)-1] == '"'
}

func IsCallbackParameter(value string) bool {
	return value[0] == '{'
}

func IsVal(value string) bool {
	return value == "val"
}
