package devices

import (
	"quando/engine/structures"
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

func After(params map[string]interface{}, ctx *structures.RunContext) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	after(duration, ctx)
}

func after(count int, ctx *structures.RunContext) {
	wait(count)
	for _, child := range ctx.CallNode.MainChildren {
		child.Method.CallFunc(ctx.Executable, child)
	}
}

func Every(params map[string]interface{}, ctx *structures.RunContext) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(count, units)
	every(duration, ctx)
}

func every(count int, ctx *structures.RunContext) {
	for {
		wait(count)
		for _, child := range ctx.CallNode.MainChildren {
			child.Method.CallFunc(ctx.Executable, child)
		}
	}
}

func Per(params map[string]interface{}, ctx *structures.RunContext) {
	count := params["count"].(int)
	units := params["units"].(string)
	duration := timeMS(1, units)
	duration = duration / count
	per(duration, ctx)
}

func per(duration int, ctx *structures.RunContext) {
	for {
		wait(duration)
		for _, child := range ctx.CallNode.MainChildren {
			child.Method.CallFunc(ctx.Executable, child)
		}
	}
}

//func ValStep(params map[string]interface{}) (float64, map[string]interface{}) {
//	val := params["val"].(float64)
//	scale := params["scale"].(float64)
//	forward := params["forward"].(bool)
//	val = valStep(forward, val, scale)
//	return val, nil
//}

func VaryOver(params map[string]interface{}, ctx *structures.RunContext) {
	count := params["count"].(int)
	units := params["units"].(string)
	mode := params["mode"].(string)
	times := params["times"].(int)
	timesUnits := params["timesUnits"].(string)
	inverted := params["inverted"].(bool)
	duration := timeMS(1, timesUnits)
	duration = duration / times
	durationTotal := timeMS(count, units)
	varyOver(duration, durationTotal, times, units, mode, inverted, ctx)

	//invert := params["invert"].(bool)
	//if invert {
	//	inverted = !inverted
	//}
	//totalTimes := times * durationTotal / 1000
	//newVal := valStep(inverted, val, totalTimes, mode)
	//switch mode {
	//case "once":
	//	if newVal > 1.0 {
	//		newVal = 1.0
	//	}
	//	if newVal < 0.0 {
	//		newVal = 0.0
	//	}
	//case "repeat":
	//	if newVal > 1.0 {
	//		newVal = 0.0
	//	}
	//	if newVal < 0.0 {
	//		newVal = 1.0
	//	}
	//case "seesaw":
	//	if newVal >= float64(1.0) || newVal <= float64(0.0) {
	//		invert = !invert
	//	}
	//}
	//data := make(map[string]interface{})
	//data["invert"] = invert
	//return newVal, data
}

func varyOver(duration int, durationTotal int, times int, units string, mode string, inverted bool, ctx *structures.RunContext) {
	totalTimes := times * durationTotal / 1000
	var val float64
	if inverted {
		val = 1.0
	} else {
		val = 0.0
	}
	for {
		wait(duration)
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
				inverted = !inverted
			}
		}
		ctx.Executable.Val = newVal
		for _, child := range ctx.CallNode.MainChildren {
			child.Method.CallFunc(ctx.Executable, child)
		}
		val = newVal
	}
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
