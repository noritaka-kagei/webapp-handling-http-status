package server

import (
	"log"
	"net/http"
	"os"
)

func Run() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	logger.Println("Setting web server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/200", OKHandler)
	logger.Println("Setting web server is completed!!")

	logger.Println("Server is running...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Printf("failed to run web server. (%s)\n", err.Error())
		logger.Println("Server is shutting down...")
		return
	}

	for {
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
