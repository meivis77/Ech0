package main

import (
    "fmt"
    "html/template"
    "net/http"
    "log"
    "golang.org/x/net/html"
)

// Variável para armazenar as rotas de cada ferramenta
type PageVariables struct {
    Title   string
    Message string
}

// Página inicial do painel
func homePage(w http.ResponseWriter, r *http.Request) {
    variables := PageVariables{
        Title:   "Painel de Controle",
        Message: "Escolha uma opção:",
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
        <ul>
            <li><a href="/scraper">Scraper</a></li>
            <li><a href="/cookie-checker">Verificador de Cookies</a></li>
            <li><a href="/link-checker">Verificador de Links Quebrados</a></li>
	    <li><a href="/html-validator">Validar HTML/CSS</a></li>
        </ul>
    </body>
    </html>`

    t, err := template.New("webpage").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}

// Manipulador para o scraper
func scraperHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        url := r.FormValue("url")
        resp, err := http.Get(url)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        doc, err := html.Parse(resp.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var links []string
        var titles []string

        var f func(*html.Node)
        f = func(n *html.Node) {
            if n.Type == html.ElementNode {
                if n.Data == "a" {
                    for _, a := range n.Attr {
                        if a.Key == "href" {
                            links = append(links, a.Val)
                        }
                    }
                }
                if n.Data == "title" {
                    titles = append(titles, n.FirstChild.Data)
                }
            }
            for c := n.FirstChild; c != nil; c = c.NextSibling {
                f(c)
            }
        }
        f(doc)

        renderScraperPage(w, links, titles)
        return
    }

    renderScraperPage(w, nil, nil)
}

// Renderiza a página do scraper
func renderScraperPage(w http.ResponseWriter, links []string, titles []string) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Scraper</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>Scraper</h1>
        <form method="POST" action="/scraper">
            <input type="text" name="url" placeholder="URL" required>
            <input type="submit" value="Scrapear">
        </form>
        {{if .}}
        <h2>Resultados:</h2>
        <h3>Títulos:</h3>
        <ul>
            {{range .Titles}}
            <li>{{.}}</li>
            {{end}}
        </ul>
        <h3>Links:</h3>
        <ul>
            {{range .Links}}
            <li><a href="{{.}}">{{.}}</a></li>
            {{end}}
        </ul>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    variables := struct {
        Links  []string
        Titles []string
    }{
        Links:  links,
        Titles: titles,
    }

    t, err := template.New("scraper").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}

// Manipulador para o verificador de cookies
func cookieCheckerHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
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

        renderCookieCheckerPage(w, cookies)
        return
    }

    renderCookieCheckerPage(w, "")
}

// Renderiza a página do verificador de cookies
func renderCookieCheckerPage(w http.ResponseWriter, cookies string) {
    tmpl := `
    <!DOCTYPE html>
    <html>
    <head>
        <title>Verificador de Cookies</title>
        <link rel="stylesheet" type="text/css" href="/static/style.css">
    </head>
    <body>
        <h1>Verificador de Cookies</h1>
        <form method="POST" action="/cookie-checker">
            <input type="text" name="url" placeholder="URL" required>
            <input type="submit" value="Verificar">
        </form>
        {{if .}}
        <h2>Cookies encontrados:</h2>
        <pre>{{.}}</pre>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    variables := struct {
        Cookies string
    }{Cookies: cookies}

    t, err := template.New("cookieChecker").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}

// Função para verificar links quebrados
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

func htmlValidatorHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        html := r.FormValue("html")
        // Aqui você pode adicionar a lógica para validar o HTML.
        // Por enquanto, vamos apenas retornar o HTML recebido.
        renderHtmlValidatorPage(w, html)
    } else {
        renderHtmlValidatorPage(w, "")
    }
}



func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/scraper", scraperHandler)
    http.HandleFunc("/cookie-checker", cookieCheckerHandler)
    http.HandleFunc("/link-checker", linkCheckerHandler)
    http.HandleFunc("/html-validator", htmlValidatorHandler) // Nova rota
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
