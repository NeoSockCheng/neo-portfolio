package handlers

import (
	"log"
	"net/http"

	"portfolio/internal/models"
)

type ProjectsData struct {
	Title    string
	Projects []models.Project
}

func (h *Handler) Projects(w http.ResponseWriter, r *http.Request) {
	projects, err := loadProjects()
	if err != nil {
		log.Printf("Error loading projects: %v", err)
		projects = []models.Project{} // Empty array if error
	}

	data := ProjectsData{
		Title:    "My Projects",
		Projects: projects,
	}

	h.renderTemplate(w, "projects.html", data)
}
