package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Printf("Starting server...")

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/health_check", check)

	http.ListenAndServe(":3000", nil)
}

// helloworld just displays a banner message for testing
func helloworld(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	w.WriteHeader(status)
	w.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head><title>Hello ` + os.Getenv("FIRSTNAME") + `</title></head>
	<body><h1>Hello ` + os.Getenv("FIRSTNAME") + `</h1></body>
	</html>
`))
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Health check</h1>`))
}
