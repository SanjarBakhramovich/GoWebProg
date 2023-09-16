package main

import (
	"GoChitChat/authenticate"
	"GoChitChat/utilities"
	"html/template"
	"net/http"
)

func main() {

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	// index
	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)
	//
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	//
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	//
	mux.HandleFunc("/authenticate", authenticate.Authenticate)
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
	//
	threads, err := data.Threads()
	if err == nil {
		_, err := utilities.Session(w, r)
		public_tmpl_files := []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}
		private_tmpl_files := []string{"templates/layout.html", "templates/public.navbar.html", "templates/index.html"}

		var templates *template.Template
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}
	}
	templates.ExecuteTemplate(w, "layout", threads)
}

// 33.5
