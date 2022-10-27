package main

import (
	"log"
	"os"

	"github.com/noritaka-kagei/webapp-handling-http-status/server4go/server"
)

func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	logger.Println("Server is starting...")
	status := server.Run()

	logger.Printf("Received quit signal: %s", status)
	logger.Println("Server is shutting down...")
}
