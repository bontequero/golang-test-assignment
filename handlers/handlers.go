package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"

	"github.com/bontequero/golang-test-assignment/models"
)

type env struct {
	models.DataLayer
}

const (
	postgresUrl = "POSTGRES_URL"
	sessionKey  = "SESSION_KEY"
	cookieName  = "auth"
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

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var reqParams loginRequest
	if err := json.NewDecoder(r.Body).Decode(&reqParams); err != nil {
		log.Printf("Cannot parse request parameters: %v; error: %v", reqParams, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := DB.GetUserInfo(reqParams.Login)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if reqParams.Password == user.Password {
		session, err := store.Get(r, cookieName)
		if err != nil {
			log.Printf("can not get cookie: %v", err)
		}

		session.Values["authenticated"] = true
		session.Save(r, w)

		w.Write([]byte("Success"))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Wrong login parameters"))
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
