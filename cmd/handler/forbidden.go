package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusForbidden(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusForbidden),
		Description:  http.StatusText(http.StatusForbidden),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusForbidden)
	w.Write(result)
}
