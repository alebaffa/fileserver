package main

import (
    "net/http"
    "os"

    "github.com/russross/blackfriday"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.Handle("/", http.FileServer(http.Dir("public")))
    http.ListenAndServe(":"+port, nil)
}