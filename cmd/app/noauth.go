package app

import (
	"gojwt/cmd/exception"
	"gojwt/cmd/handler"
	"net/http"
)

func noAuth(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()
	handler.StatusOk(w, M{
		"message": "this is no auth",
	})
}
