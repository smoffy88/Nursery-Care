package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World from your Go web app on Project IDX!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server listening...")
    http.ListenAndServe(":8080", nil)
}
