package devices

import (
	"quando/engine/structures"
	"time"
)

func GetTimerCallbacks() []structures.Method {
	return []structures.Method{
		structures.Method{
			Identifier: "quando.time.after",
			Function:   After,
			Type:       "callback",
			Iterator:   false,
			Arbiter:    false,
		},
		structures.Method{
			Identifier: "quando.time.every",
			Function:   Every,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
		structures.Method{
			Identifier: "quando.time.per",
			Function:   Per,
			Type:       "callback",
			Iterator:   true,
			Arbiter:    false,
		},
	}
}

func After(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	after(duration)
	return 0.0, nil
}

func Every(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	after(duration)
	return 0.0, nil
}

func Per(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(1, units)
	duration = duration / count
	after(duration)
	return 0.0, nil
}

func after(count int) {
	time.Sleep(time.Duration(count) * time.Millisecond)
}

func timeMS(value int, unit string) int {
	switch unit {
	case "milliseconds":
		return value
	case "seconds":
		return value * 1000
	case "minutes":
		return value * 1000 * 60
	case "hours":
		return value * 1000 * 60 * 60
	case "days":
		return value * 1000 * 60 * 60 * 24
	default:
		return value
	}
}
