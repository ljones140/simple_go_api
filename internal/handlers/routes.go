package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ljones140/simple_go_api/internal/handlers/object"
	"github.com/ljones140/simple_go_api/internal/repository"
)

func RegisterRoutes(mux *chi.Mux, repo repository.Repository) {
	ob := object.New(repo)
	mux.Route("/objects", func(r chi.Router) {
		r.Post("/", ob.Create)
		r.Get("/{id}", ob.Get)
	})
}
