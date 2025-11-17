package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"portfolio/internal/config"
	"portfolio/internal/handlers"
)

//go:embed ../../web/templates/* ../../web/static/*
var embeddedFiles embed.FS

func main() {
	// Load configuration
	cfg := config.New()

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// Static file server
	staticFS := http.FS(embeddedFiles)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	// Initialize handlers
	h := handlers.New(embeddedFiles)

	// Routes
	r.Get("/", h.Home)
	r.Get("/about", h.About)
	r.Get("/projects", h.Projects)
	r.Get("/contact", h.Contact)
	r.Post("/contact", h.ContactSubmit)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}

	_ = staticFS // Keep for production builds
}
