package controller

import (
	"encoding/json"
	// "fmt"
	"gorestful/model"
	"net/http"
)

func All(w http.ResponseWriter, r *http.Request) {
	response := model.GetUserAll()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
