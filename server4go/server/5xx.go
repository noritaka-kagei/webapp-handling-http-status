/*
 * This file handling HTTP response status based URI parameter
 * Target HTTP response status: Server error responses (500 - 599)
 */

package server

import "net/http"

func Set5XX() {
	// index of HTTP status 500~511
	http.HandleFunc("/5XX", fiveHundredHandler)
	http.HandleFunc("/5xx", fiveHundredHandler)

	// 500 ~ 511
	http.HandleFunc("/500", internalServerErrorHandler)
	http.HandleFunc("/501", notImplementedHandler)
	http.HandleFunc("/502", badGatewayHandler)
	http.HandleFunc("/503", serviceUnavailableHandler)
	http.HandleFunc("/504", gatewayTimeoutHandler)
	http.HandleFunc("/505", httpVersionNotSupportedHandler)
	http.HandleFunc("/506", variantAlsoNegotiatesHandler)
	http.HandleFunc("/507", insufficientStorageHandler)
	http.HandleFunc("/508", loopDetectedHandler)
	http.HandleFunc("/510", notExtendedHandler)
	http.HandleFunc("/511", networkAuthenticationRequiredHandler)
}

func fiveHundredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := "You can access following HTTP status (5XX)\n"
	content += "500, 501, 502, 503, 504, 505, 506, 507, 508,\n"
	content += "510, 511\n"
	w.Write([]byte(content))
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
}

func badGatewayHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(502)
}

func serviceUnavailableHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(503)
}

func gatewayTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(504)
}

func httpVersionNotSupportedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(505)
}

func variantAlsoNegotiatesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(506)
}

func insufficientStorageHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(507)
}

func loopDetectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(508)
}

func notExtendedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(510)
}

func networkAuthenticationRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(511)
}
