package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var host = "localhost"
var port = 8080

func Run() os.Signal {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	logger.Println("Setting web server...")
	http.HandleFunc("/", rootHandler)
	Set1XX()
	Set2XX()
	Set3XX()
	Set4XX()
	Set5XX()
	logger.Println("Setting web server is completed!!")

	logger.Println("Server is running...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil); err != nil {
		logger.Printf("failed to run web server. (%s)\n", err.Error())
		logger.Println("Server is shutting down...")
		return nil
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	return <-quit
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := fmt.Sprintf("You can get response with HTTP status XXX when you access to \"%s:%d/XXX\" via HTTP protocol.\n", host, port)
	content += fmt.Sprintf("You can show avaiable path(XXX) by /NXX that prefix of hundreds place. (ex. %s:%d/2XX)\n", host, port)
	w.Write([]byte(content))
}
