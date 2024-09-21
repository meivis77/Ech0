# Ech0 v.000.002

Esta ferramenta inovadora foi programada por um bot com o objetivo de se tornar uma solução abrangente para desenvolvedores e profissionais da web. Após várias atualizações e melhorias, esperamos que ela se transforme em uma ferramenta de grande ajuda, oferecendo uma variedade de funções úteis para diferentes necessidades. (Isso tudo é apenas um experimento)

# ![Banner](Screenshot_20240920-203533.png) <!-- Substitua pela URL da sua imagem de banner -->

# Tecnologia Utilizada

Desenvolvida em **Go**, a ferramenta é executada diretamente no navegador, proporcionando uma experiência rápida e responsiva. Ao contrário de outras ferramentas disponíveis, nossa abordagem única combina eficiência e facilidade de uso, destacando-se no mercado.


1. **Web Scraper:**

Extrai informações relevantes de páginas da web, como links, usuários e títulos. Essa funcionalidade é crucial para aqueles que precisam coletar dados rapidamente para análise ou desenvolvimento.



2. **Cookie Checker:**

Permite a extração de informações do cookie de uma página da web. Com isso, os usuários podem obter todos os detalhes relevantes sobre os cookies, auxiliando em auditorias de segurança e conformidade.



3. **Verificador de Links Quebrados:**

Auxilia na verificação de vulnerabilidades ao identificar links quebrados em páginas. Essa função é essencial para manter a integridade e a usabilidade de sites, ajudando a evitar problemas que possam afetar a experiência do usuário.



4. **Validação de HTML/CSS:**

Oferece suporte no estudo da estrutura web, permitindo que os desenvolvedores validem seu código HTML e CSS. Esta funcionalidade não apenas ajuda a garantir que os sites sejam semanticamente corretos, mas também a melhorar a acessibilidade e o SEO.

---

## Funcionalidades

- Scraper web (Extrai informações de páginas web link, users,title) **Auxilia na injeção sqlmap**
- Cookie Checker (Extrai informações do cookie de uma página web) **Todas as informações**
- Verificador de Links Quebrados **Auxilia na verificação de vulnerabilidades**
- Validar HTML/CSS **Ajuda no estudo de estrutura web**

---

## Pré-requisitos

Talvez você precise instalar algumas dependências se algo não rodar bem.

- Abaixando (**opcional**):

1. **Instalar o Go**: [Baixe aqui](https://golang.org/dl/).
2. **Instalar o Node.js**: [Baixe aqui](https://nodejs.org/).
3. **Após instalar o Node.js instale o html-validator
```bash
   npm install -g html-validator
```

- Pelo Terminal (**recomendado**)

1. **Atualize os pack's**:
```bash
pkg upgrade && pkg update
```
2. **Instale Golang**:
```bash
pkg install golang
```
3. **Instale Node.js**:
```bash
pkg install nodejs
```

- Dependência da Ferramenta :
```bash
npm install -g html-validator
```
**and**
```bash
go mod tidy
```

### Comando para Rodar a Ferramenta:

- Dentro do diretório você rodará a ferramenta com os seguintes comandos:

```bash
go build painel.gp
```
```bash
./painel.go
```
**Solução Alternativa:**
```bash
go run painel.go
```

# Ferramenta Web:

- Para ter acesso as suas ferramentas,basta entrar no site :

```bash
http://localhost:8080/
```

# Estrutura dos arquivos (Suporte)

```bash
ech0/
│
├── convert.txt              # Arquivo de texto para conversão
├── go.mod                   # Arquivo de módulo Go, contém dependências do projeto
├── go.sum                   # Arquivo de soma, mantém um registro das versões das dependências
│
├── convert_to_txt           # Diretório ou arquivo (precisa de mais contexto)
├── cookie_checker            # Diretório ou arquivo (precisa de mais contexto)
├── painel                    # Diretório ou arquivo (precisa de mais contexto)
├── scraper.go               # Código fonte do scraper
├── html-validator.go         # Código fonte do validador HTML
├── link_checker.go           # Código fonte do verificador de links
├── painel.go                 # Código fonte principal do painel
├── cookie_checker.go         # Código fonte do verificador de cookies
└── servidor.go               # Código fonte do servidor
└── static                    # Diretório para arquivos estáticos (CSS, JS, etc.)
```
