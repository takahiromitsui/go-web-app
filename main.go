package main

import (
	"fmt"
	"net/http"
)

const port = ":8000"

func main() {
	fs := http.FileServer(http.Dir("src/public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, nil)
}

// ├── main.go
// ├── tailwind.config.js
// └── src
//     ├── styles.css(input.css)
//     └── public
//         ├── home.html
//         └── styles.css (output.css)
