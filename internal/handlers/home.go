package handlers

import (
	"net/http"
)

type HomeData struct {
	Title       string
	Description string
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Title:       "Home",
		Description: "Welcome to my portfolio",
	}

	h.renderTemplate(w, "home.html", data)
}
