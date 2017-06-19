package main

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var signUp = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	var user string

	err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

	switch {
	case err == sql.ErrNoRows:
		var hashedPassword []byte
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}
		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			http.Error(w, "Server error, unable to create your account.", 500)
			return
		}
		w.Write([]byte("User created!"))
		return

	case err != nil:
		http.Error(w, "Server error, unable to create your account.", 500)
		return

	case err == nil:
		http.Error(w, "User already exists.", 409)
		return

	default:
	}
})
