package main

import (
    "bytes"
    "html/template"
    "net/http"
    "os/exec"
)

// Manipulador para o validador de HTML e CSS
func htmlValidatorHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        htmlContent := r.FormValue("html")
        cmd := exec.Command("html-validator", "--verbose", "--format", "text", "--", "-")
        cmd.Stdin = bytes.NewBufferString(htmlContent)

        output, err := cmd.CombinedOutput()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        renderHtmlValidatorPage(w, string(output))
        return
    }

    renderHtmlValidatorPage(w, "")
}

// Renderiza a página do validador de HTML e CSS
func renderHtmlValidatorPage(w http.ResponseWriter, result string) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Validador de HTML e CSS</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>Validador de HTML e CSS</h1>
        <form method="POST" action="/html-validator">
            <textarea name="html" placeholder="Cole seu HTML aqui..." required></textarea>
            <input type="submit" value="Validar">
        </form>
        {{if .Result}}
        <h2>Resultados:</h2>
        <pre>{{.Result}}</pre>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    variables := struct {
        Result string
    }{Result: result}

    t, err := template.New("htmlValidator").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, variables)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/html-validator", htmlValidatorHandler)
    http.ListenAndServe(":8080", nil)
}
