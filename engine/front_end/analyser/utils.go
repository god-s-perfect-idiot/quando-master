package analyser

import (
	"strconv"
	"strings"
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

func FindLineType(line []string) string {
	if IsLineBreak(line) {
		return "lineBreak"
	} else if IsInvocation(line) {
		return "invocation"
	} else if IsCallbackTerminator(line) {
		return "callbackTerminator"
	} else if IsConditionalCallback(line) {
		return "conditionalCallback"
	} else {
		return "unknown"
	}
}

func IsLineBreak(line []string) bool {
	if len(line) == 0 {
		return true
	}
	return false
}

func IsInvocation(line []string) bool {
	if strings.Contains(line[0], "quando.") {
		return true
	}
	return false
}

func IsCallbackTerminator(line []string) bool {
	if len(line) == 1 && line[0] == "}" {
		return true
	}
	return false
}

func IsConditionalCallback(line []string) bool {
	if len(line) == 3 && line[0] == "}" && line[1] == "callback" && line[2] == "{" {
		return true
	}
	return false
}
