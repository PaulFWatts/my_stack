package web

import (
	"embed"
	"html/template"
	"net/http"
	"time"
)

//go:embed templates/*.html
var templateFiles embed.FS

type Server struct {
	templates *template.Template
}

type PageData struct {
	Title string
	Year  int
}

func NewServer() (*Server, error) {
	templates, err := template.ParseFS(templateFiles, "templates/*.html")
	if err != nil {
		return nil, err
	}

	return &Server{templates: templates}, nil
}

func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.index)
	mux.HandleFunc("/greeting", s.greeting)

	return mux
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	data := PageData{
		Title: "Go + HTMX + Tailwind Demo",
		Year:  time.Now().Year(),
	}

	if err := s.templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) greeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "greeting.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
