package main

import (
    "fmt"
    "net/http"
)

// Verifica se um link está quebrado
func checkLink(url string) (bool, error) {
    resp, err := http.Head(url)
    if err != nil {
        return false, err
    }
    return resp.StatusCode < 400, nil
}

// Manipulador para o verificador de links quebrados
func linkCheckerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        url := r.FormValue("url")
        isAlive, err := checkLink(url)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var message string
        if isAlive {
            message = fmt.Sprintf("O link %s está funcionando.", url)
        } else {
            message = fmt.Sprintf("O link %s está quebrado.", url)
        }

        renderLinkCheckerPage(w, message)
        return
    }

    renderLinkCheckerPage(w, "")
}

// Renderiza a página do verificador de links quebrados
func renderLinkCheckerPage(w http.ResponseWriter, message string) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Verificador de Links Quebrados</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>Verificador de Links Quebrados</h1>
        <form method="POST" action="/link-checker">
            <input type="text" name="url" placeholder="URL" required>
            <input type="submit" value="Verificar">
        </form>
        {{if .}}
        <h2>Resultado:</h2>
        <p>{{.}}</p>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    variables := struct {
        Message string
    }{Message: message}

    t, err := template.New("linkChecker").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}
