package app

import (
	"gojwt/cmd/config"
	"gojwt/cmd/exception"
	"gojwt/cmd/handler"
	"gojwt/cmd/service"
	"net/http"
)

type M map[string]interface{}

func login(w http.ResponseWriter, r *http.Request) {
	defer exception.CatchUp()

	if r.Method == http.MethodPost {
		username, password, ok := r.BasicAuth()
		if !ok {
			handler.StatusBadRequest(w, M{
				"message": "bad request for credential",
			})

			return
		}

		user := config.GetConfig()
		for _, v := range user.Users {
			isValid := v.Username == username && v.Password == password

			if isValid {
				token, err := service.GenerateToken(v.Username)
				if err != nil {
					handler.StatusInternalServerError(w, M{
						"message": "Internal server error",
					})
				}

				handler.StatusOk(w, M{
					"message": "success",
					"token":   token,
				})

				return
			}
		}
	}

	handler.StatusUnauthorized(w, M{"message": "username or password is wrong"})
}
