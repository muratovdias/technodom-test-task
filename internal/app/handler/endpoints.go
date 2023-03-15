package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/muratovdias/technodom-test-task/internal/app/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/admin", func(r chi.Router) {
		r.Get("/redirects", h.allLinks)
		r.Get("/redirects/{id}", h.getLink)
		r.Patch("/redirects/{id}", h.updateLink)
		r.Delete("/redirects/{id}", h.deleteLink)
		r.Post("/redirects", h.createLink)
	})
	r.Get("/redirects", h.redirect)
	return r
}
