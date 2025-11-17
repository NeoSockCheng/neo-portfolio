package api

import (
	"embed"
	"encoding/json"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"sync"
)

//go:embed ../web/templates/**/*.html
var templatesFS embed.FS

//go:embed ../web/static/**/*
var staticFS embed.FS

//go:embed ../data/projects.json
var projectsData []byte

var (
	router    http.Handler
	once      sync.Once
	templates *template.Template
)

type Project struct {
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Technologies []string `json:"technologies"`
	GithubURL    string   `json:"github_url"`
	LiveURL      string   `json:"live_url"`
	ImageURL     string   `json:"image_url"`
}

func initHandler() {
	once.Do(func() {
		// Parse templates from embedded filesystem
		var err error
		templates, err = template.ParseFS(templatesFS, "web/templates/layouts/*.html", "web/templates/pages/*.html", "web/templates/partials/*.html")
		if err != nil {
			log.Printf("Error parsing templates: %v", err)
		}

		// Setup router
		mux := http.NewServeMux()

		// Static files from embedded filesystem
		staticFiles, err := fs.Sub(staticFS, "web/static")
		if err != nil {
			log.Printf("Error setting up static files: %v", err)
		}
		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))

		// Routes
		mux.HandleFunc("/", homeHandler)
		mux.HandleFunc("/about", aboutHandler)
		mux.HandleFunc("/projects", projectsHandler)
		mux.HandleFunc("/contact", contactHandler)

		router = mux
	})
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	if templates == nil {
		http.Error(w, "Templates not loaded", http.StatusInternalServerError)
		return
	}
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Template error: %v", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Title": "Home",
	}
	renderTemplate(w, "home.html", data)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "About Me",
	}
	renderTemplate(w, "about.html", data)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// Parse projects from embedded data
	var projects []Project
	if err := json.Unmarshal(projectsData, &projects); err != nil {
		log.Printf("Error parsing projects JSON: %v", err)
	}

	data := map[string]interface{}{
		"Title":    "Projects",
		"Projects": projects,
	}
	renderTemplate(w, "projects.html", data)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle form submission
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		log.Printf("Contact form submission - Name: %s, Email: %s, Message: %s", name, email, message)

		data := map[string]interface{}{
			"Title":   "Contact",
			"Success": true,
			"Message": "Thank you for your message! I'll get back to you soon.",
		}
		renderTemplate(w, "contact.html", data)
		return
	}

	data := map[string]interface{}{
		"Title": "Contact",
	}
	renderTemplate(w, "contact.html", data)
}

// Handler is the entry point for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	initHandler()
	router.ServeHTTP(w, r)
}
