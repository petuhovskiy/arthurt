package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-type", "text/html; charset=utf-8")
		fmt.Fprintln(w, registerHTML)
		err := r.URL.Query().Get("error")
		if err != "" {
			fmt.Fprintln(w, `<div style="border: 1px solid red">`+html.EscapeString(err)+`</div>`)
		}
	case "POST":
		processRegister(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func processRegister(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("user")
	password := r.Form.Get("password")     // TODO: validation
	_, err := findUserByUsername(username) // TODO: mongo unique indexes
	if err != mgo.ErrNotFound {
		http.Redirect(w, r, "/register?error=User+already+exists", http.StatusFound)
		return
	}
	user := &User{
		ID:       bson.NewObjectId(),
		Username: username,
		Password: password,
	}
	err = insertNewUser(user)
	log.Printf("%#v\n", user)
	if err != nil {
		http.Redirect(w, r, "/register?error="+url.QueryEscape(err.Error()), http.StatusFound)
		return
	}
	log.Printf("User %s registered successfully", username)
	writeSession(w, jwt.MapClaims{
		"UserId": user.ID.Hex(),
	})
	http.Redirect(w, r, "/", http.StatusFound)
}
