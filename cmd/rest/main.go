package main

import (
	"os"
	"os/signal"

	"62tech.co/service/config/env"
	"62tech.co/service/handler"
)

func main() {
	env.LoadEnv()

	// Init dependencies
	service := handler.InitHandler()

	// start echo server
	service.StartServer()

	// Shutdown with gracefull handler
	service.ShutdownServer()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
