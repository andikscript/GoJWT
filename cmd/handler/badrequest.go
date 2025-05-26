package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusBadRequest(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusBadRequest),
		Description:  http.StatusText(http.StatusBadRequest),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(result)
}

func BadRequestList(w http.ResponseWriter, maps []map[string]interface{}) {
	response := model.ResponseList{
		ResponseCode: fmt.Sprintf("%d", http.StatusBadRequest),
		Description:  http.StatusText(http.StatusBadRequest),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(result)
}
