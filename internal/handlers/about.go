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
			"Swift", "Objective-C", "UIKit", "SwiftUI",
			"Golang", "Python", "Java", "PHP",
			"JavaScript", "Dart", "HCL",
			"React", "Flutter", "Codeigniter",
			"MySQL", "Firebase",
			"AWS", "Terraform", "Docker",
			"Git", "XCode", "Power BI", "Figma",
		},
	}

	h.renderTemplate(w, "about.html", data)
}
