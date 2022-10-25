package server

import "net/http"

func OKHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
