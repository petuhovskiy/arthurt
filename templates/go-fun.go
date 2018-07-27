package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/petuhovskiy/go-fun/wiki"
)

func main() {
	fmt.Print("Starting application...\n")
	t1 := new(int)
	t2 := t1

	fmt.Println(*t1, *t2)

	*t1 += 2
	*t2 -= 2

	fmt.Println(*t1, *t2)

	log.Print("Ok")

	l1 := 2

	y1 := &l1

	fmt.Println(l1, *y1)

	*y1 += 3
	l1--

	fmt.Println(l1, *y1)

	wiki.InitHandlers()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
