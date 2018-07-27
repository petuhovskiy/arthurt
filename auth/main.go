package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	port    = 9090
	portStr = strconv.Itoa(port)
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page!")
	userID, ok := authStatus(r)
	if ok {
		fmt.Fprintf(w, "You are logged in as %s\n", userID)
	} else {
		fmt.Fprintln(w, "You are not logged in")
	}
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)

	log.Printf("Listen on port %d\n", port)
	err := http.ListenAndServe(":"+portStr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
