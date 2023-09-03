package app

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"quando/engine/back_end/core"
	"quando/engine/back_end/generator"
	"quando/engine/front_end/analyser"
	"quando/engine/structures"
)

var callPipe *chan map[string]interface{}
var memory *structures.Memory

func runScript(script string, callPipe *chan map[string]interface{}) {
	hash := structures.GetHash(script)
	var executable *structures.Executable
	if memory.Check(hash) {
		executable = memory.Get(hash).(*structures.Executable)
	} else {
		codeAnalyser := analyser.NewAnalyser(script)
		executable = codeAnalyser.Scan()
		executable.ConnectCallPipe(callPipe)
		codeGenerator := generator.NewGenerator()
		codeGenerator.GenerateCode(*executable)
		memory.Set(hash, executable)
	}
	core.Execute(executable)
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

func crashScript(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		r.ParseForm()
		script := r.Form.Get("script")
		fmt.Println("Received POST request with body:", script)
		if script != "" {
			hashID := structures.GetHash(script)
			*callPipe <- map[string]interface{}{
				"crash": hashID,
			}
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Script Termination Scheduled"))
	}
}

func Listen(channel *chan map[string]interface{}) {
	mux := http.NewServeMux()
	memory = structures.NewMemory()
	mux.HandleFunc("/script", readScript)
	mux.HandleFunc("/stop", crashScript)
	callPipe = channel
	fmt.Println("Quando Go Engine started")
	fmt.Println("..serving Engine on : 127.0.0.1:1024")
	handler := cors.Default().Handler(mux)
	err := http.ListenAndServe(":1024", handler)
	if err != nil {
		panic(err)
	}
}
