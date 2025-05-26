package service

import (
	"context"
	"fmt"
	"gojwt/cmd/handler"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var APPLICATION_NAME = "Go JWT"
var LOGIN_EXPIRED_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("secret-from-go-jwt-token")

func MiddlewareJwtAuth(next http.Handler) http.Handler {

	noAuth := []string{"/login", "/noauth"}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range noAuth {
			if r.URL.Path == v {
				next.ServeHTTP(w, r)
				return
			}
		}

		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			handler.StatusBadRequest(w, map[string]interface{}{"message": "Invalid token"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		tokenString = strings.TrimSpace(tokenString)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if method, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signning method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})

		if err != nil {
			handler.StatusBadRequest(w, map[string]interface{}{"message": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			handler.StatusBadRequest(w, map[string]interface{}{"message": "Invalid token"})
			return
		}

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
