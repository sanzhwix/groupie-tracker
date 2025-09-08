package internal

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var Tpl *template.Template

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if 	r.Method != http.MethodGet {
		RenderErrorPage(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	artistStruct, err := FetchArtist()
	if err != nil {
		log.Fatal("Error fetching atist in main")
	}
	Tpl.ExecuteTemplate(w, "index.html", artistStruct)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
}

type ErrorMsg struct {
	Msg string
}
func RenderErrorPage(w http.ResponseWriter, msg string, statusCode int) {
	w.WriteHeader(statusCode)
	data := ErrorMsg{Msg:msg}
	err := Tpl.ExecuteTemplate(w,"errorPage.html", data)
	if err != nil {
		fmt.Println(" Error with executing Error Page template!")
		return 
	}
}