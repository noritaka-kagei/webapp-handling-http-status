package server

import "net/http"

func Set2XX() {
	// 200 ~ 226
	http.HandleFunc("/200", okHandler)
	http.HandleFunc("/201", createdHandler)
	http.HandleFunc("/202", acceptedHandler)
	http.HandleFunc("/203", nonAuthInfoHandler)
	http.HandleFunc("/204", noContentHandler)
	http.HandleFunc("/205", resetContentHandler)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func createdHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}

func acceptedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(202)
}

func nonAuthInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(203)
}

func noContentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func resetContentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(205)
}

// TODO: 206

// ToDo: 207

// ToDo: 208

// ToDo: 226
