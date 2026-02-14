# 🚀 RapidPDF

> **The fastest way to turn the web into paper (virtually)!** / **O jeito mais rápido de transformar a web em papel (virtualmente)!**

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)

---

### 🌐 Select Language / Selecione o Idioma

- [🇧🇷 Português Brasileiro (PT-BR)](#-português-brasileiro)
- [🇺🇸 English (EN)](#-english)

---

## 🇧🇷 Português Brasileiro

### 🤔 O que é isso?

Cansado de salvar páginas da web uma por uma? O **RapidPDF** é sua ferramenta CLI (Linha de Comando) mágica que pega um monte de links, renderiza tudo como se fosse um navegador de verdade (obrigado, Chrome! 🤖) e costura tudo num único arquivo PDF bonitão.

É tipo um álbum de figurinhas da internet, só que útil.

### ✨ O que ele faz?

- **Multitarefa**: Passou 1, 2 ou 10 URLs? Ele engole tudo.
- **Modo Servidor**: Rode sem argumentos e ele vira uma API REST pronta para o combate. 🛡️
- **Documentado**: Swagger UI incluído de fábrica. 🎩
- **Inteligente**: Usa o motor do Chrome (via `chromedp`) para garantir que o PDF fique igualzinho ao site.
- **Organizado**: Junta (merge) todas as páginas em um arquivo `output.pdf` final.
- **Seguro**: Valida suas URLs para você não digitar besteira.
- **Configurável**: Tem limites para você não exagerar (controlado via `.env`).

### 🚀 Bora rodar

#### 1. Modo CLI (Clássico)

Mande as URLs e veja a mágica acontecer:

```bash
go run main.go https://go.dev https://google.com
```

_Boom!_ 💥 Veja o arquivo `output.pdf` aparecer na sua pasta.

#### 2. Modo Servidor (API Power)

Rode sem argumentos para subir o servidor:

```bash
go run main.go
# 🚀 RapidPDF — Web-to-PDF Converter
# 📡 Server listening on :8080
```

Agora você tem superpoderes:

- **Gerar PDF**: `POST /generate` com JSON `{"urls": ["..."]}`
- **Documentação**: Acesse [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) e brinque com a API.

### 🛠️ Tecnologias (O Motor)

Debaixo do capô, temos monstros sagrados do ecossistema Go:

- 🐹 **Go**: Porque gostamos de velocidade.
- 🍸 **Gin**: O framework web mais rápido do oeste.
- 📜 **Swagger**: Documentação automática para ninguém ficar perdido.
- 🌐 **Chromedp**: Para renderizar as páginas com precisão cirúrgica.
- 📄 **pdfcpu**: Para colar os PDFs uns nos outros sem usar cola tenaz.
- 📝 **godotenv**: Porque hardcoded config é coisa do passado.

---

## 🇺🇸 English

### 🤔 What is this?

Tired of saving web pages one by one? **RapidPDF** is your magical tool that takes a bunch of links, renders them just like a real browser (thanks, Chrome! 🤖), and stitches them all together into a single, beautiful PDF file.

Run it as a CLI or start it as a REST API server. You choose!

### ✨ What does it do?

- **Multitasking**: 1, 2, or 10 URLs? It eats them for breakfast.
- **Server Mode**: Run without args to start a robust REST API. 🛡️
- **Documented**: Swagger UI included out of the box. 🎩
- **Smart**: Uses the Chrome engine (via `chromedp`) to ensure the PDF looks exactly like the website.
- **Organized**: Merges everything into a final `output.pdf` file.
- **Safe**: Validates your URLs so you don't type nonsense.
- **Configurable**: Has limits so you don't go overboard (controlled via `.env`).

### 🚀 Let's run it

#### 1. CLI Mode (Classic)

Feed it URLs and watch it fly:

```bash
go run main.go https://go.dev https://google.com
```

_Boom!_ 💥 Watch the `output.pdf` file appear in your folder.

#### 2. Server Mode (API Power)

Run without arguments to launch the server:

```bash
go run main.go
# 🚀 RapidPDF — Web-to-PDF Converter
# 📡 Server listening on :8080
```

Now you have superpowers:

- **Generate PDF**: `POST /generate` with JSON `{"urls": ["..."]}`
- **Documentation**: Go to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) and play with the API.

### 🛠️ Tech Stack (The Engine)

Under the hood, we have some heavy hitters from the Go ecosystem:

- 🐹 **Go**: Because we like speed.
- 🍸 **Gin**: The fastest web framework in the wild west.
- 📜 **Swagger**: Automatic docs so you never get lost.
- 🌐 **Chromedp**: To render pages with surgical precision.
- 📄 **pdfcpu**: To glue PDFs together without using actual glue.
- 📝 **godotenv**: Because hardcoded config is so last season.

---

### 📜 License

MIT © [Paulo Silva](https://github.com/psilva1982)

_Made with ❤️ and a lot of caffeine._
