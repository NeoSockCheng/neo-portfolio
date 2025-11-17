package handlers

import (
	"log"
	"net/http"
)

type ContactData struct {
	Title   string
	Message string
	Success bool
}

func (h *Handler) Contact(w http.ResponseWriter, r *http.Request) {
	data := ContactData{
		Title: "Contact Me",
	}

	h.renderTemplate(w, "contact.html", data)
}

func (h *Handler) ContactSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// TODO: Implement email sending logic here
	// For now, just log the submission
	log.Printf("Contact form submission - Name: %s, Email: %s, Message: %s", name, email, message)

	data := ContactData{
		Title:   "Contact Me",
		Message: "Thank you for your message! I'll get back to you soon.",
		Success: true,
	}

	h.renderTemplate(w, "contact.html", data)
}
