package main

import (
    "html/template"
    "net/http"
    "path"
    "os"
    "math/rand"
)

type Teammate struct {
    Name  string
}

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
    teammate := []Teammate{
        {"Alessandro"},
        {"Davide"},
        {"Marcello"},
        {"Sarah"},
        {"Fred"},
        {"Khaled"},
        {"Herve"},
        {"Bo"},
        {"Remi"},
    }

    fp := path.Join("public", "winner.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(rw, teammate[rand.Intn(len(teammate))]); err != nil {
        http.Error(rw, err.Error(), http.StatusInternalServerError)
    }
}