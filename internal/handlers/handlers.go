package handlers

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"portfolio/internal/models"
)

type Handler struct {
	templates *template.Template
}

func New(embeddedFiles embed.FS) *Handler {
	// Parse all templates
	tmpl, err := template.ParseFS(embeddedFiles, "web/templates/**/*.html")
	if err != nil {
		// Fallback to local files for development
		tmpl, err = template.ParseGlob("web/templates/**/*.html")
		if err != nil {
			log.Fatal("Error parsing templates:", err)
		}
	}

	return &Handler{
		templates: tmpl,
	}
}

func (h *Handler) renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	err := h.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func loadProjects() ([]models.Project, error) {
	data, err := os.ReadFile("data/projects.json")
	if err != nil {
		return nil, err
	}

	var projects []models.Project
	err = json.Unmarshal(data, &projects)
	return projects, err
}

// Helper to check if templates are embedded
func (h *Handler) isEmbedded(embeddedFiles embed.FS) bool {
	_, err := fs.Stat(embeddedFiles, "web/templates")
	return err == nil
}
