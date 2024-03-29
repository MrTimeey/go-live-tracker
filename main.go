package main

import (
	"fmt"

	"html/template"
	"log"
	"net/http"
)

func main() {
	templateFunction := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templates/index.html"))
		tmpl.Execute(w, nil)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", templateFunction)

	fmt.Println("Server up and running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
