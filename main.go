package main

import (
	"fmt"
	"net/http"
	engine "quando/engine/app"
	"quando/internal/config"
	"quando/internal/server"
	"quando/internal/server/ip"
)

var handlers = []server.Handler{} // extra handlers are added when running full version

func fetchHandler(route string) server.Handler {
	for _, handler := range handlers {
		if handler.Url == route {
			return handler
		}
	}
	return server.Handler{}
}

func runServer() {
	fmt.Println("Quando Go Server started")
	ipAddress := ip.PrivateIP()
	if config.RemoteClient() {
		fmt.Println("  Client can be accessed remotely at ", ipAddress)
	}
	if config.RemoteEditor() {
		fmt.Println("**SECURITY WARNING** Editor can be accessed remotely at ", ipAddress)
	}
	server.ServeHTTPandIO(handlers)
}

func main() {

	callPipe := make(chan map[string]interface{})

	length := len(handlers)
	println("handlers length: ", length)
	go engine.Listen(&callPipe)
	runServer()

	fmt.Println("Go listener running for engine -> server in main thread")

	mux := http.NewServeMux()

	for _, handler := range handlers {
		mux.HandleFunc(handler.Url, handler.Func)
	}

	//for {
	//	select {
	//	case <-callPipe:
	//		//route := callData["route"].(string)
	//		//println("Received call: ", callData["route"].(string), callData["body"])
	//		//println("Received")
	//		//handler := fetchHandler(route)
	//		//println("Handler: ", handler.Url)
	//		println("I was probably not expecting this")
	//	}
	//}
}
