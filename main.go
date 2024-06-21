package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", index)
	r.Get("/user-info", user)
	r.Get("/task/{id}", id)

	server := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]string)
	ctx["Name"] = "Ben"
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, ctx)

	if err != nil {
		log.Println("Template Execution Error")
	}
}
func user(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}
func id(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Hello! %s", id)
}
