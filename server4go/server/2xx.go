package server

import "net/http"

func Set2XX() {
	// index of HTTP status 200~226
	http.HandleFunc("/2XX", twoHundredHandler)
	http.HandleFunc("/2xx", twoHundredHandler)

	// 200 ~ 226
	http.HandleFunc("/200", okHandler)
	http.HandleFunc("/201", createdHandler)
	http.HandleFunc("/202", acceptedHandler)
	http.HandleFunc("/203", nonAuthInfoHandler)
	http.HandleFunc("/204", noContentHandler)
	http.HandleFunc("/205", resetContentHandler)
	http.HandleFunc("/206", partialContentHandler)
	http.HandleFunc("/207", multiStatusHandler)
	http.HandleFunc("/208", alreadyReportedHandler)
	http.HandleFunc("/226", imUsedHandler)
}

func twoHundredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := "You can access following HTTP status (2XX)\n"
	content += "200, 201, 202, 203, 204, 205, 206, 207, 208, 226\n"
	w.Write([]byte(content))
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

func partialContentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(206)
}

func multiStatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(207)
}

func alreadyReportedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(208)
}

func imUsedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(226)
}
