package generator

import (
	"quando/engine/back_end/actions"
	"quando/engine/devices"
	"quando/engine/front_end/analyser"
	"quando/engine/structures"
)

type Lookups struct {
	Actions   *structures.LookupTable
	Callbacks *structures.LookupTable
}

func GenerateLookups() *Lookups {
	apiClient := actions.NewClient()
	deviceClient := devices.NewClient()
	lookups := &Lookups{
		Actions:   structures.ConstructLookupTable(apiClient.GetActions()),
		Callbacks: structures.ConstructLookupTable(deviceClient.GetCallbacks()),
	}
	return lookups
}

func Run() {
	testQuery := `
	quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
		quando.control.key ch="a", up_down="down", on_off=val
	}
	quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
		quando.control.key ch="a", up_down="down", on_off=val
	}
	quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
		quando.control.key ch="a", up_down="down", on_off=val
	}
	`
	codeAnalyser := analyser.NewAnalyser(testQuery)
	essence := codeAnalyser.Scan()
	println(essence)
	generator := NewGenerator()
	generator.GenerateCode(*essence)
}
