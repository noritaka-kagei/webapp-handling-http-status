/*
 * This file handling HTTP response status based URI parameter
 * Target HTTP response status: Informational responses (100 - 199)
 */

package server

import "net/http"

func Set1XX() {
	// index of HTTP status 100~103
	http.HandleFunc("/1XX", oneHundredHandler)
	http.HandleFunc("/1xx", oneHundredHandler)

	// 100 ~ 103
	http.HandleFunc("/100", continueHandler)
	http.HandleFunc("/101", switchingProtocolsHandler)
	http.HandleFunc("/102", processingHandler)
	http.HandleFunc("/103", earlyHintsHandler)
}

func oneHundredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := "You can access following HTTP status (1XX)\n"
	content += "100, 101, 102, 103\n"
	w.Write([]byte(content))
}

func continueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(100)
}

func switchingProtocolsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(101)
}

func processingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(102)
}

func earlyHintsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(103)
}
