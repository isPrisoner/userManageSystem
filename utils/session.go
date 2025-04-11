package utils

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"net/url"
)

func SessionInit(w http.ResponseWriter, username string) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return
	}
	sessionID := url.QueryEscape(base64.URLEncoding.EncodeToString(b))
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})
	SessionStoreDB(sessionID, username)
}

func Session(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
	})
}

func SessionStoreDB(sessionID, username string) {
	db := InitDB()
	stmt, err := db.Prepare("update user set session_id = ? where username = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(sessionID, username)
}

func SessionDelete(w http.ResponseWriter, sessionID string) {
	db := InitDB()
	stmt, err := db.Prepare("update user set session_id = NULL where session_id = ?")
	defer stmt.Close()
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(sessionID)
	http.SetCookie(w, &http.Cookie{
		Name:     "sessionID",
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}
