package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/internal"
)

func main() {
	var err error
	internal.Tpl, err = template.ParseGlob("./internal/templates/*.html")
	if err != nil {
		log.Fatal("Something went wrong while parsing templates!")
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./internal/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", internal.MainHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
