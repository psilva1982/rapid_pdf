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
- **Inteligente**: Usa o motor do Chrome (via `chromedp`) para garantir que o PDF fique igualzinho ao site.
- **Organizado**: Junta (merge) todas as páginas em um arquivo `output.pdf` final.
- **Seguro**: Valida suas URLs para você não digitar besteira.
- **Configurável**: Tem limites para você não exagerar (controlado via `.env`).

### 🚀 Bora rodar

1.  **Instale as dependências** (certifique-se de ter o Go instalado):

    ```bash
    go mod download
    ```

2.  **Configure o ambiente**:
    Crie um arquivo `.env` (se não tiver) e defina o limite de URLs:

    ```env
    MAX_URLS=10
    ```

3.  **Execute a mágica**:

    ```bash
    go run main.go https://go.dev https://google.com
    ```

    _Boom!_ 💥 Veja o arquivo `output.pdf` aparecer na sua pasta.

### 🛠️ Tecnologias (O Motor)

Debaixo do capô, temos monstros sagrados do ecossistema Go:

- 🐹 **Go**: Porque gostamos de velocidade.
- 🌐 **Chromedp**: Para renderizar as páginas com precisão cirúrgica.
- 📄 **pdfcpu**: Para colar os PDFs uns nos outros sem usar cola tenaz.
- 📝 **godotenv**: Porque hardcoded config é coisa do passado.

---

## 🇺🇸 English

### 🤔 What is this?

Tired of saving web pages one by one? **RapidPDF** is your magical CLI (Command Line Interface) tool that takes a bunch of links, renders them just like a real browser (thanks, Chrome! 🤖), and stitches them all together into a single, beautiful PDF file.

It's like a sticker album for the internet, but actually useful.

### ✨ What does it do?

- **Multitasking**: 1, 2, or 10 URLs? It eats them for breakfast.
- **Smart**: Uses the Chrome engine (via `chromedp`) to ensure the PDF looks exactly like the website.
- **Organized**: Merges everything into a final `output.pdf` file.
- **Safe**: Validates your URLs so you don't type nonsense.
- **Configurable**: Has limits so you don't go overboard (controlled via `.env`).

### 🚀 Let's run it

1.  **Install dependencies** (make sure you have Go installed):

    ```bash
    go mod download
    ```

2.  **Configure the environment**:
    Create a `.env` file (if you don't have one) and set the URL limit:

    ```env
    MAX_URLS=10
    ```

3.  **Do the magic**:

    ```bash
    go run main.go https://go.dev https://google.com
    ```

    _Boom!_ 💥 Watch the `output.pdf` file appear in your folder.

### 🛠️ Tech Stack (The Engine)

Under the hood, we have some heavy hitters from the Go ecosystem:

- 🐹 **Go**: Because we like speed.
- 🌐 **Chromedp**: To render pages with surgical precision.
- 📄 **pdfcpu**: To glue PDFs together without using actual glue.
- 📝 **godotenv**: Because hardcoded config is so last season.

---

### 📜 License

MIT © [Paulo Silva](https://github.com/psilva1982)

_Made with ❤️ and a lot of caffeine._
