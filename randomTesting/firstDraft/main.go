package main

import (
    "fmt"
    "net/http"
    "log"
    "html/template"
    "strings"
)

func login(w http.ResponseWriter,r *http.Request){
    fmt.Println("method: ",r.Method)

    if r.Method == "GET" {
        t, _ = template.ParseFiles("./frontend/login.html")
    } else {
    }
}

func main() {
    
    // http.HandleFunc("/",frontpage)
    http.HandleFunc("/login", login)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

