package handlers

import (
	"net/http"
)

type AboutData struct {
	Title string
	Skills []string
}

func (h *Handler) About(w http.ResponseWriter, r *http.Request) {
	data := AboutData{
		Title: "About Me",
		Skills: []string{
			"Go",
			"JavaScript",
			"HTML/CSS",
			"Git",
		},
	}

	h.renderTemplate(w, "about.html", data)
}
