package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", "Welcome!")
	fmt.Fprintf(w, "<h2>%s</h2>", "Web application written in Golang")
}

func healthchkHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm fine")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	body := "Just a wiki page, Lorem ipsum dolor sit amet, consectetur adipiscing elit."
    fmt.Fprintf(w, "<b>%s</b>", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthcheck/", healthchkHandler)
	http.HandleFunc("/wiki/", viewHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
