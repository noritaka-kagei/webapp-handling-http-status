/*
 * This file handling HTTP response status based URI parameter
 * Target HTTP response status: Redirection messages (300 - 399)
 */

package server

import "net/http"

func Set3XX() {
	// index of HTTP status 300~308
	http.HandleFunc("/3XX", threeHundredHandler)
	http.HandleFunc("/3xx", threeHundredHandler)

	// 300 ~ 308
	http.HandleFunc("/300", multipleChoicesHandler)
	http.HandleFunc("/301", movedPermanentlyHandler)
	http.HandleFunc("/302", foundHandler)
	http.HandleFunc("/303", seeOtherHandler)
	http.HandleFunc("/304", notModifiedHandler)
	http.HandleFunc("/305", useProxyHandler)
	http.HandleFunc("/306", unusedHandler)
	http.HandleFunc("/307", temporaryRedirectHandler)
	http.HandleFunc("/308", permanentRedirectHandler)
}

func threeHundredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := "You can access following HTTP status (3XX)\n"
	content += "300, 301, 302, 303, 304, 305, 306, 307, 308\n"
	w.Write([]byte(content))
}

func multipleChoicesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(300)
}

func movedPermanentlyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(301)
}

func foundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(302)
}

func seeOtherHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(303)
}

func notModifiedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(304)
}

func useProxyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(305)
}

func unusedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(306)
}

func temporaryRedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(307)
}

func permanentRedirectHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(308)
}
