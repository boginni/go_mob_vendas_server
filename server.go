package main

import (
	"net/http"
)

type response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Params  string `json:"params"`
}

var (
	host = "192.168.1.3:11005"
)

func serve() error {	
	return http.ListenAndServe(host, getRoutes())
}
