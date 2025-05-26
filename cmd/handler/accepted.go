package handler

import (
	"encoding/json"
	"fmt"
	"gojwt/cmd/model"
	"net/http"
)

func StatusAccepted(w http.ResponseWriter, maps map[string]interface{}) {
	response := model.Response{
		ResponseCode: fmt.Sprintf("%d", http.StatusAccepted),
		Description:  http.StatusText(http.StatusAccepted),
		Data:         maps,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result, _ := json.Marshal(response)
	w.WriteHeader(http.StatusAccepted)
	w.Write(result)
}
