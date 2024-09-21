package main

import (
    "fmt"
    "html/template"
    "net/http"
    "log"
)

type PageVariables struct {
    Title   string
    Message string
    Cookies string
}

func homePage(w http.ResponseWriter, r *http.Request) {
    variables := PageVariables{
        Title:   "Verificador de Cookies",
        Message: "Insira uma URL para verificar os cookies:",
    }

    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>{{.Title}}</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>{{.Message}}</h1>
        <form method="POST" action="/check-cookies">
            <input type="text" name="url" placeholder="URL" required>
            <input type="submit" value="Verificar">
        </form>
        {{if .Cookies}}
        <h2>Cookies encontrados:</h2>
        <pre>{{.Cookies}}</pre>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    t, err := template.New("webpage").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}

func checkCookiesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    url := r.FormValue("url")
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var cookies string
    for _, cookie := range resp.Cookies() {
        cookies += fmt.Sprintf("Nome: %s, Valor: %s, Domínio: %s, Expiração: %s\n",
            cookie.Name, cookie.Value, cookie.Domain, cookie.Expires)
    }

    variables := PageVariables{
        Title:   "Verificador de Cookies",
        Message: "Insira uma URL para verificar os cookies:",
        Cookies: cookies,
    }

    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>{{.Title}}</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>{{.Message}}</h1>
        <form method="POST" action="/check-cookies">
            <input type="text" name="url" placeholder="URL" required>
            <input type="submit" value="Verificar">
        </form>
        {{if .Cookies}}
        <h2>Cookies encontrados:</h2>
        <pre>{{.Cookies}}</pre>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    t, err := template.New("webpage").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/check-cookies", checkCookiesHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
    log.Fatal(http.ListenAndServe(":8082", nil))
}
