package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusCreated(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusCreated),
		Description:  http.StatusText(http.StatusCreated),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}
