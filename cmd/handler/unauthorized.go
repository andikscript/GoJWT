package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusUnauthorized(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusUnauthorized),
		Description:  http.StatusText(http.StatusUnauthorized),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(result)
}
