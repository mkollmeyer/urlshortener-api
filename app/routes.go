package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mkollmeyer/urlshortener-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/url", loadUrlRoutes)
	return router
}

func loadUrlRoutes(router chi.Router) {
	urlHandler := &handler.Url{}
	router.Post("/", urlHandler.Gen)
	router.Get("/{short}", urlHandler.Get)
}
