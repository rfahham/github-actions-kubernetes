package main

import (
    "fmt"
    "net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "<h1>Pipeline All in One GitHub Actions!!!</h1>")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}
