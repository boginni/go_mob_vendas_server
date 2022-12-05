package routes

import "net/http"


type EchoHandler struct {
	defaultResponse string
}

func NewEchoHandler(defaultResponse string) *EchoHandler {
	return &EchoHandler{defaultResponse: defaultResponse}
}

func (e EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(e.defaultResponse))
	w.WriteHeader(http.StatusOK)
}
