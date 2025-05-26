package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusMethodNotAllowed(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusMethodNotAllowed),
		Description:  http.StatusText(http.StatusMethodNotAllowed),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write(result)
}
