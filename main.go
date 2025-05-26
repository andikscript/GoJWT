package main

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/app"
	"gojwt/cmd/config"
	"gojwt/cmd/service"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func init() {
	if config.SharedConfig != nil {
		return
	}

	basePath, err := os.ReadFile(filepath.Join("assets", "./user.json"))
	if err != nil {
		log.Fatal("Error read file json", err.Error())
		return
	}

	err = json.Unmarshal(basePath, &config.SharedConfig)
	if err != nil {
		log.Fatal("Error conversion,", err.Error())
		return
	}
}

func main() {
	mux := &service.CustomMux{}

	// jwt
	mux.RegisterMiddleware(service.MiddlewareJwtAuth)

	// route
	app.Router(mux)

	address := fmt.Sprintf(":%d", 8123)
	server := &http.Server{
		Addr:         address,
		Handler:      mux,
		ReadTimeout:  40 * time.Second,
		WriteTimeout: 40 * time.Second,
	}

	log.Println("⇨ Listening and serving HTTPS on", address)

	err := server.ListenAndServeTLS("./cert/server.crt", "./cert/server.key")
	if err != nil {
		log.Fatal("Error when start and listen server -", err.Error())
		return
	}
}
