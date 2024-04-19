package main

import (
	"fmt"
	"net/http"

	"github.com/takahiromitsui/go-web-app/pkg/handlers"
)

const port = ":8000"

func main() {
	fs := http.FileServer(http.Dir("src/public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}

