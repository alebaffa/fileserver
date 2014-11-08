package main

import (
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    //http.Handle("/", http.FileServer(http.Dir("public")))
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public"))))
    http.HandleFunc("/assign", AssignPTR)
    http.ListenAndServe(":"+port, nil)
}

func AssignPTR(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
}