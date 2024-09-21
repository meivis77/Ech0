package main

import (
    "html/template"
    "net/http"
    "net/http/cookiejar"
    "strings"
    "golang.org/x/net/html"
)

func scraperHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        url := r.FormValue("url")
        jar, _ := cookiejar.New(nil)
        client := &http.Client{Jar: jar}

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        resp, err := client.Do(req)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            http.Error(w, "Erro ao acessar a URL: "+resp.Status, http.StatusInternalServerError)
            return
        }

        doc, err := html.Parse(resp.Body)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var sb strings.Builder
        var f func(*html.Node)
        f = func(n *html.Node) {
            if n.Type == html.ElementNode && n.Data == "a" {
                for _, attr := range n.Attr {
                    if attr.Key == "href" {
                        sb.WriteString(n.FirstChild.Data + ": " + attr.Val + "\n")
                    }
                }
            }
            for c := n.FirstChild; c != nil; c = c.NextSibling {
                f(c)
            }
        }
        f(doc)

        results := sb.String()
        renderScraperPage(w, results)
        return
    }

    renderScraperPage(w, "")
}

func renderScraperPage(w http.ResponseWriter, results string) {
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
            <label for="url">URL:</label>
            <input type="text" id="url" name="url" required>
            <br>
            <input type="submit" value="Scrape">
        </form>
        {{if .Results}}
        <h2>Resultados:</h2>
        <pre>{{.Results}}</pre>
        {{end}}
        <a href="/">Voltar ao Painel</a>
    </body>
    </html>`

    variables := struct {
        Results string
    }{Results: results}

    t, err := template.New("scraper").Parse(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, variables)
}
