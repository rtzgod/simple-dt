package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	})
	r.Mount("/data", h.dataRoutes())

	return r
}

func (h *Handler) dataRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.SetData)
	r.Get("/", h.GetData)

	return r
}
