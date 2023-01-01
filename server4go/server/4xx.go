package server

import "net/http"

func Set4XX() {
	// index of HTTP status 400~451
	http.HandleFunc("/4XX", fourHundredHandler)
	http.HandleFunc("/4xx", fourHundredHandler)

	// 400 ~ 451
	http.HandleFunc("/400", badRequestHandler)
	http.HandleFunc("/401", unauthorizedHandler)
	http.HandleFunc("/402", paymentRequiredHandler)
	http.HandleFunc("/403", forbiddenHandler)
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/405", methodNotAllowedHandler)
	http.HandleFunc("/406", notAcceptableHandler)
	http.HandleFunc("/407", proxyAuthenticationRequiredHandler)
	http.HandleFunc("/408", requestTimeoutHandler)
	http.HandleFunc("/409", conflictHandler)
	http.HandleFunc("/410", goneHandler)
	http.HandleFunc("/411", lengthRequiredHandler)
	http.HandleFunc("/412", preconditionFailedHandler)
	http.HandleFunc("/413", payloadTooLargeHandler)
	http.HandleFunc("/414", uriTooLongHandler)
	http.HandleFunc("/415", unsupportedMediaTyoeHandler)
	http.HandleFunc("/416", rangeNotSatisfitableHandler)
	http.HandleFunc("/417", expectationFailedHandler)
	http.HandleFunc("/418", imaTeapotHandler)
	http.HandleFunc("/421", misdirectedRequestHandler)
	http.HandleFunc("/422", unprocessableEntityHandler)
	http.HandleFunc("/423", lockedHandler)
	http.HandleFunc("/424", failedDependencyHandler)
	http.HandleFunc("/425", tooEarlyHandler)
	http.HandleFunc("/426", upgradeRequiredHandler)
	http.HandleFunc("/428", preconditionRequiredHandler)
	http.HandleFunc("/429", tooManyRequestsHandler)
	http.HandleFunc("/431", requestHeaderFieldsTooLargeHandler)
	http.HandleFunc("/451", unavailableForLegalReasonsHandler)
}

func fourHundredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	content := "You can access following HTTP status (4XX)\n"
	content += "400, 401, 402, 403, 404, 405, 406, 407, 408, 409,\n"
	content += "410, 411, 412, 413, 414, 415, 416, 417, 418,\n"
	content += "421, 422, 423, 424, 425, 426, 428, 429,\n"
	content += "431, 451\n"
	w.Write([]byte(content))
}

func badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(400)
}

func unauthorizedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
}

func paymentRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(402)
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(403)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
}

func notAcceptableHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(406)
}

func proxyAuthenticationRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(407)
}

func requestTimeoutHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(408)
}

func conflictHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(409)
}

func goneHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(410)
}

func lengthRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(411)
}

func preconditionFailedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(412)
}

func payloadTooLargeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(413)
}

func uriTooLongHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(414)
}

func unsupportedMediaTyoeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(415)
}

func rangeNotSatisfitableHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(416)
}

func expectationFailedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(417)
}

func imaTeapotHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(418)
}

func misdirectedRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(421)
}

func unprocessableEntityHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(422)
}

func lockedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(423)
}

func failedDependencyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(424)
}

func tooEarlyHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(425)
}

func upgradeRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(426)
}

func preconditionRequiredHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(428)
}

func tooManyRequestsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(429)
}

func requestHeaderFieldsTooLargeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(431)
}

func unavailableForLegalReasonsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(451)
}
