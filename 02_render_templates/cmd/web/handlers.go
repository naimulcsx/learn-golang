package main

import (
	"html/template"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Add("Server", "Go")

	files := []string{
		"ui/html/base.layout.tmpl",
		"ui/html/partials/nav.partial.tmpl",
		"ui/html/partials/footer.partial.tmpl",
		"ui/html/pages/home.page.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"ui/html/base.layout.tmpl",
		"ui/html/partials/nav.partial.tmpl",
		"ui/html/partials/footer.partial.tmpl",
		"ui/html/pages/about.page.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	id := r.PathValue("id")

	data := struct {
		UserID string
	}{
		UserID: id,
	}

	files := []string{
		"ui/html/base.layout.tmpl",
		"ui/html/partials/nav.partial.tmpl",
		"ui/html/partials/footer.partial.tmpl",
		"ui/html/pages/user.page.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base.layout.tmpl", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"running","version":"1.0.0"}`))
}