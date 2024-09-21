package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Olá, mundo!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor rodando em http://localhost:8080/")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Erro ao iniciar o servidor:", err)
    }
}
