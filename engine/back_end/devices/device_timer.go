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
		structures.Method{
			Identifier: "quando.time.vary",
			Function:   VaryOver,
			Type:       "callback",
			DataKeys:   []string{"invert"},
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
	return -1.0, nil
}

func Every(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	after(duration)
	return -1.0, nil
}

func Per(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(1, units)
	duration = duration / count
	after(duration)
	return -1.0, nil
}

//func ValStep(params map[string]interface{}) (float64, map[string]interface{}) {
//	val := params["val"].(float64)
//	scale := params["scale"].(float64)
//	forward := params["forward"].(bool)
//	val = valStep(forward, val, scale)
//	return val, nil
//}

func VaryOver(params map[string]interface{}) (float64, map[string]interface{}) {
	count := params["count"].(int)
	units := params["units"].(string)
	mode := params["mode"].(string)
	times := params["times"].(int)
	timesUnits := params["timesUnits"].(string)
	inverted := params["inverted"].(bool)
	val := params["val"].(float64)
	duration := timeMS(1, timesUnits)
	duration = duration / times
	durationTotal := timeMS(count, units)
	if params["invert"] == nil {
		params["invert"] = false
		val = 0.0
	}
	invert := params["invert"].(bool)
	if invert {
		inverted = !inverted
	}
	totalTimes := times * durationTotal / 1000
	newVal := valStep(inverted, val, totalTimes, mode)
	switch mode {
	case "once":
		if newVal > 1.0 {
			newVal = 1.0
		}
		if newVal < 0.0 {
			newVal = 0.0
		}
	case "repeat":
		if newVal > 1.0 {
			newVal = 0.0
		}
		if newVal < 0.0 {
			newVal = 1.0
		}
	case "seesaw":
		if newVal >= float64(1.0) || newVal <= float64(0.0) {
			invert = !invert
		}
	}
	data := make(map[string]interface{})
	data["invert"] = invert
	after(duration)
	return newVal, data
}

func valStep(inverted bool, val float64, times int, mode string) float64 {
	scale := 1.0 / float64(times)
	if !inverted {
		val += scale
	} else {
		val -= scale
	}
	switch mode {
	case "once":
		if val > 1.0 {
			val = 1.0
		}
		if val < 0.0 {
			val = 0.0
		}
	case "repeat":
		if val > 1.0 {
			val = 0.0
		}
		if val < 0.0 {
			val = 1.0
		}
	}

	return val
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
