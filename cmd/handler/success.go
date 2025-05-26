package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusOk(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusOK),
		Description:  http.StatusText(http.StatusOK),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func StatusOkList(w http.ResponseWriter, maps []map[string]interface{}) {
	response := model.ResponseList{
		ResponseCode: fmt.Sprintf("%d", http.StatusOK),
		Description:  http.StatusText(http.StatusOK),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
