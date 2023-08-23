package core

import (
	"quando/engine/back_end/actions"
	"quando/engine/back_end/generator"
	"quando/engine/devices"
	"quando/engine/front_end/analyser"
)

type Client struct {
	apiClient    *actions.Client
	deviceClient *devices.Client
}

func NewClient() *Client {
	apiClient := actions.NewClient()
	deviceClient := devices.NewClient()
	client := &Client{
		apiClient:    apiClient,
		deviceClient: deviceClient,
	}
	return client
}

func Run() {
	//client := NewClient()

	//testQuery := `
	//quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
	//	quando.control.key ch="a", up_down="down", on_off=val
	//}
	//quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
	//	quando.control.key ch="a", up_down="down", on_off=val
	//}
	//quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
	//	quando.control.key ch="a", up_down="down", on_off=val
	//}
	//`
	testQuery := `
	quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
		quando.key.handleKey id=0, key="a", ctrl=false, alt=false, callback = {
			quando.control.key ch="b", upDown="down", onOff=val
		}	
		quando.control.key ch="b", upDown="down", onOff=val
	}
	`
	codeAnalyser := analyser.NewAnalyser(testQuery)
	essence := codeAnalyser.Scan()
	codeGenerator := generator.NewGenerator()
	codeGenerator.GenerateCode(*essence)
	Execute(essence)
}
