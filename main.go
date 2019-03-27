package main

import (
	"html/template"
	"net/http"
	"log"
)

var TPL *template.Template


func init() {
	TPL = template.Must(template.ParseGlob("*.html"))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
  	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", index)
	http.HandleFunc("/solution", solution)
	//http.HandleFunc("/client", client)
	//http.HandleFunc("/about", about)
	//http.HandleFunc("/contact", contact)

	log.Println("Listening...")
	http.ListenAndServe(":7600", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	data := struct {

		Title string

	}{
		Title: "Home",
		
	}
	TPL.ExecuteTemplate(w, "index.html", data)

}


func solution(w http.ResponseWriter, r *http.Request) {

	data := struct {

		Title string

	}{
		Title: "Solution",
		
	}
	TPL.ExecuteTemplate(w, "solution.html", data)

}





