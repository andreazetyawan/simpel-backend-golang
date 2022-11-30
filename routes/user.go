package user

import(
	"net/http"
	"os"
	
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func user() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	return r
}