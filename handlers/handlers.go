package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", login)

	r.Route("/api", func(r chi.Router) {
		r.Route("/notes", func(r chi.Router) {
			r.Post("/add", addNote)
			r.Get("/", getAllNotes)

			r.Route("/{noteID}", func(r chi.Router) {
				r.Get("/", getNote)
				r.Delete("/", deleteNote)
			})
		})
	})

	return r
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func addNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
