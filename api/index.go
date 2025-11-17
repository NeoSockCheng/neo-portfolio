package api

import (
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"portfolio/internal/handlers"
)

var (
	router *chi.Mux
	once   sync.Once
)

func initRouter() {
	once.Do(func() {
		// Initialize router
		router = chi.NewRouter()

		// Middleware
		router.Use(middleware.Logger)
		router.Use(middleware.Recoverer)
		router.Use(middleware.Compress(5))

		// Static file server
		router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

		// Initialize handlers
		h := handlers.New()

		// Routes
		router.Get("/", h.Home)
		router.Get("/about", h.About)
		router.Get("/projects", h.Projects)
		router.Get("/contact", h.Contact)
		router.Post("/contact", h.ContactSubmit)
	})
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	initRouter()
	router.ServeHTTP(w, r)
}
