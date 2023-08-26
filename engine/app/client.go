package app

import (
	"fmt"
	"net/http"
	"quando/engine/back_end/core"
	"quando/engine/back_end/generator"
	"quando/engine/front_end/analyser"
)

var callPipe *chan map[string]interface{}

func runScript(script string, callPipe *chan map[string]interface{}) {
	codeAnalyser := analyser.NewAnalyser(script)
	essence := codeAnalyser.Scan()
	essence.ConnectCallPipe(callPipe)
	codeGenerator := generator.NewGenerator()
	codeGenerator.GenerateCode(*essence)
	core.Execute(essence)
}

func readScript(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		script := r.Form.Get("script")
		fmt.Println("Received POST request with body:", script)
		if script != "" {
			go runScript(script, callPipe)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Script received"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Listen(channel *chan map[string]interface{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/script", readScript)
	callPipe = channel
	fmt.Println("Quando Go Engine started")
	fmt.Println("..serving Engine on : 127.0.0.1:1024")
	err := http.ListenAndServe(":1024", mux)
	if err != nil {
		panic(err)
	}
}
