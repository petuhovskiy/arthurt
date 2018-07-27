package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("superSecretToken$112223333")
var sessionCookie = "GO_SESSION"
var jwtSigningMethod = jwt.SigningMethodHS256

func writeSession(w http.ResponseWriter, claims jwt.MapClaims) {
	cur := time.Now().UTC()
	exp := cur.Add(time.Hour * 5)
	claims["nbf"] = cur.Unix()
	claims["exp"] = exp.Unix()
	token := jwt.NewWithClaims(jwtSigningMethod, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Printf("JWT Token Error: %s", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   sessionCookie,
		Value:  tokenString,
		MaxAge: 60 * 60 * 5,
	})
}

func readSession(r *http.Request) (jwt.MapClaims, error) {
	cookie, err := r.Cookie(sessionCookie)
	if err != nil {
		return nil, err
	}
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("JWT token invalid")
	}
}

func authStatus(r *http.Request) (id string, ok bool) {
	claims, err := readSession(r)
	if err != nil {
		ok = false
		return
	}
	id, ok = claims["UserId"].(string)
	return
}
