package handler

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Value string `json:"value"`
}

var data Data

func (h *Handler) SetData(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetData(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
