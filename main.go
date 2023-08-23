package main

import (
	"fmt"
	engine "quando/engine/app"
	"quando/internal/config"
	"quando/internal/server"
	"quando/internal/server/ip"
)

var handlers = []server.Handler{} // extra handlers are added when running full version

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

func runEngine() {
	callPipe := make(chan string)
	go engine.Listen(&callPipe)

	for {
		select {
		case call := <-callPipe:
			println("Received call: ", call)
		}
	}

}

func main() {
	go runEngine()
	runServer()
}
