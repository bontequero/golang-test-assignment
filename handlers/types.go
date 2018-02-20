package handlers

import (
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"

	"github.com/bontequero/golang-test-assignment/models"
)

type env struct {
	models.DataLayer
}

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type noteData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

const (
	postgresUrl = "POSTGRES_URL"

	sessionKey   = "SESSION_KEY"
	cookieName   = "auth-cookie"
	cookieAuth   = "authenticated"
	cookieRole   = "role"
	cookieUserID = "user-id"
)

var (
	DB        *env
	secretKey = []byte(os.Getenv(sessionKey))
	store     = sessions.NewCookieStore(secretKey)
)

func NewEnv() (*env, error) {
	db, err := models.NewDB(os.Getenv(postgresUrl))
	if err != nil {
		return nil, err
	}

	DB = &env{db}
	return DB, nil
}

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", login)
	r.Post("/logout", logout)

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
