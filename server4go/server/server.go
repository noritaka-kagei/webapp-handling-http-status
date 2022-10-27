package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() os.Signal {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	logger.Println("Setting web server...")
	http.HandleFunc("/", rootHandler)
	Set2XX()
	logger.Println("Setting web server is completed!!")

	logger.Println("Server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
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
}
