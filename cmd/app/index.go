package app

import (
	"gojwt/cmd/exception"
	"gojwt/cmd/handler"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func index(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	claims, ok := r.Context().Value("userInfo").(jwt.MapClaims)
	if !ok {
		handler.StatusUnauthorized(w, M{
			"message": "Unauthorized",
		})

		return
	}

	log.Println("claims:", claims)

	handler.StatusOk(w, M{
		"message": "this is index",
	})
}
