package main

import (
	"net/http"
	"os"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/andreazetyawan/routes"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3333"
	} 
	port = ":" + port

	return port
}

func main(){
	f := routes.user()
	// r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world"))
	// })

	http.ListenAndServe(getPort(), f)
}