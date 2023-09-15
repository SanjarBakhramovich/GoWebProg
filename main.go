package main

import (
	"html/template"
	"net/http"
)

func main() {

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	// index
	mux.Handler("/", index)
	mux.HandleFunc("/err", err)
	//
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	//
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	//
	mux.HandleFunc("/authenticate", authenticate)
	//
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	//
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/post", readThread)
	//
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
