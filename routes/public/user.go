package public

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type authActionRoutine func(w http.ResponseWriter, r *http.Request) error

type AuthHandler struct {
	dbConn     interface{}
	authAction string
}

func NewAuthHandler(authAction string) *AuthHandler {
	return &AuthHandler{dbConn: nil, authAction: authAction}
}

func (authH AuthHandler) getEndpointHandler() authActionRoutine {
	handler, ok := map[string]authActionRoutine{
		"validate": authH.validate,
		"login":    authH.login,
	}[authH.authAction]

	if !ok {
		return func(w http.ResponseWriter, r *http.Request) error {
			w.WriteHeader(http.StatusNotFound)
			return errors.New("not existent endpoint")
		}
	}
	return handler
}

func (authH *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := authH.getEndpointHandler()
	if err := handler(w, r); err != nil {
		fmt.Printf("err: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (authH *AuthHandler) validate(w http.ResponseWriter, r *http.Request) error {
	token := r.Header.Get("token")

	if strings.Compare(token, "acb") != 0 {
		message := "invalid session token"
		w.Write([]byte(message))
		return errors.New(message)
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func (authH *AuthHandler) login(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte(`{"token": "abc"}`))
	return err
}
