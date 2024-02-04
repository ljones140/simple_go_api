package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ljones140/simple_go_api/internal/handlers/object"
)

func RegisterRoutes(mux *chi.Mux) {
	ob := object.New()
	mux.Route("/objects", func(r chi.Router) {
		r.Post("/", ob.Create)
	})
}
