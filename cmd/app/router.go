package app

import (
	"gojwt/cmd/exception"
	"gojwt/cmd/service"
)

func Router(mux *service.CustomMux) {
	defer exception.CatchUp()

	mux.HandleFunc("/login", login)

	mux.HandleFunc("/index", index)

	mux.HandleFunc("/noauth", noAuth)
}
