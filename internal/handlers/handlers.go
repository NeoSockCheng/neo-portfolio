package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"portfolio/internal/models"
)

type Handler struct {
	templates *template.Template
}

func New() *Handler {
	// Parse all templates from disk
	tmpl := template.New("")

	// Parse all template files
	layouts, _ := filepath.Glob("web/templates/layouts/*.html")
	pages, _ := filepath.Glob("web/templates/pages/*.html")
	partials, _ := filepath.Glob("web/templates/partials/*.html")

	allFiles := append(layouts, pages...)
	allFiles = append(allFiles, partials...)

	if len(allFiles) == 0 {
		log.Fatal("No template files found")
	}

	tmpl, err := tmpl.ParseFiles(allFiles...)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
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
