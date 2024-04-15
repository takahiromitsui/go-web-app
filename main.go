package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8000"

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.page.html"))
	tmpl.Execute(w, nil)
}

// func About(w http.ResponseWriter, r *http.Request) {
	
// }

func main() {
	http.HandleFunc("/", Home)
	// http.HandleFunc("/about", About)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}

