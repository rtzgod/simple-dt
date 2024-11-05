package handler

import (
	"encoding/json"
	"github.com/rtzgod/simple-dt/internal/entity"
	"net/http"
)

var data entity.Data

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
