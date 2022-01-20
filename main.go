package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	d, e := os.Getwd()
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(d, "    :8080")

	http.Handle("/", http.FileServer(http.Dir(d)))
	e = http.ListenAndServe(":8080", nil)
	if e != nil {
		log.Fatal(e)
	}
}
