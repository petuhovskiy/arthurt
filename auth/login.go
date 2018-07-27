package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"

	"github.com/dgrijalva/jwt-go"
)

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-type", "text/html; charset=utf-8")
		fmt.Fprintln(w, loginHTML)
		err := r.URL.Query().Get("error")
		if err != "" {
			fmt.Fprintln(w, `<div style="border: 1px solid red">`+html.EscapeString(err)+`</div>`)
		}
	case "POST":
		processLogin(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func processLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("user")
	password := r.Form.Get("password")
	user, err := findUserByUsername(username)
	if err != nil {
		http.Redirect(w, r, "/login?error="+url.QueryEscape(err.Error()), http.StatusFound)
		return
	}
	if user.Password != password {
		http.Redirect(w, r, "/login?error=Incorrect+login+or+password", http.StatusFound)
		return
	}
	log.Printf("User %s, login success\n", username)
	writeSession(w, jwt.MapClaims{
		"UserId": user.ID.Hex(),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
