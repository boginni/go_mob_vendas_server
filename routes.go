package main

import (
	"net/http"

	"github.com/boginni/go_mob_vendas_server/routes"
	"github.com/boginni/go_mob_vendas_server/routes/middlewares"
	public_routes "github.com/boginni/go_mob_vendas_server/routes/public"
	"github.com/gorilla/mux"
)

func getRoutes() http.Handler {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()

	public := v1.PathPrefix("/public").Subrouter()
	public.Use(
		middlewares.DDOS,
	)

	// GET host/v1/public/login
	user := public.PathPrefix("/user").Subrouter()

	user.Handle("/login", public_routes.NewAuthHandler("login")).Methods(http.MethodGet)
	user.Handle("/validate", public_routes.NewAuthHandler("validate")).Methods(http.MethodGet)

	echo := v1.PathPrefix("/echo").Subrouter()

	echo.Use(
		middlewares.DDOS,
	)

	echo.Handle("/1", routes.NewEchoHandler(`{"status": "OK"}`)).Methods(http.MethodGet)
	echo.Handle("/2", routes.NewEchoHandler(`{"status": "OK"}`)).Methods(http.MethodPost)

	return r
}
