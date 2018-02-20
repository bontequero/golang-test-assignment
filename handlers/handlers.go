package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

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

		session.Values[cookieAuth] = true
		session.Values[cookieRole] = user.Role
		session.Values[cookieUserID] = user.ID
		session.Save(r, w)

		w.Write([]byte("Success"))
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Wrong login parameters"))
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, cookieName)
	if err != nil {
		log.Printf("can not get cookie: %v", err)
	}

	session.Values[cookieAuth] = false
	session.Values[cookieRole] = ""
	session.Values[cookieUserID] = 0
	session.Save(r, w)

	w.WriteHeader(http.StatusOK)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, cookieName)
	if err != nil {
		log.Printf("can not get cookie: %v", err)
		http.Error(w, "Cookie is invalid", http.StatusBadRequest)
		return
	}

	if auth, ok := session.Values[cookieAuth].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	var reqParams noteData
	if err := json.NewDecoder(r.Body).Decode(&reqParams); err != nil {
		log.Printf("Cannot parse request parameters: %v; error: %v", reqParams, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DB.AddNote(map[string]interface{}{
		"userID":  session.Values[cookieUserID].(int64),
		"title":   reqParams.Title,
		"content": reqParams.Content,
	})
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Added"))
}

func getAllNotes(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, cookieName)
	if err != nil {
		log.Printf("can not get cookie: %v", err)
		http.Error(w, "Cookie is invalid", http.StatusBadRequest)
		return
	}

	if auth, ok := session.Values[cookieAuth].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	notes, err := DB.GetAllNotes(
		session.Values[cookieUserID].(int64),
		session.Values[cookieRole].(string),
	)
	if err != nil {
		log.Print(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(notes); err != nil {
		log.Printf("can not encode response from db: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}
}

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(""))
}
