# 🚀 RapidPDF

> **Turning the wild wild web into pristine PDF paper!** / **Transformando a web selvagem em papel PDF impecável!**

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue?style=flat-square)](LICENSE)

---

### 🌐 Select Language / Selecione o Idioma

- [🇧🇷 Português Brasileiro (PT-BR)](#-português-brasileiro)
- [🇺🇸 English (EN)](#-english)

---

## 🇧🇷 Português Brasileiro

### 🤔 O que é isso?

Cansado de "Salvar como PDF" página por página? O **RapidPDF** é sua varinha mágica via CLI (Linha de Comando) que engole URLs, renderiza com a precisão de um navegador real (valeu, Chrome! 🤖) e costura tudo num único PDF bonitão. É tipo um álbum de figurinhas da internet, só que útil.

### ✨ Superpoderes

- **Multitarefa**: 1, 10 ou 100 URLs? Manda ver.
- **Modo Servidor**: Rode sem argumentos e ele vira uma API REST poderosa. 🛡️
- **Documentado**: Swagger UI incluído, porque ninguém merece adivinhar rotas. 🎩
- **Inteligente**: Usa o motor do Chrome (`chromedp`) para garantir que o PDF fique _igualzinho_ ao site.
- **Organizado**: Junta (merge) todas as páginas em um arquivo `output.pdf` final.
- **Seguro**: Valida suas URLs para você não passar vergonha.
- **Configurável**: Limites de URLs e Timeout ajustáveis via `.env` (porque o tempo é dinheiro).

### 🚀 Bora rodar

#### 1. Modo CLI (Clássico)

Mande as URLs e veja a mágica acontecer:

```bash
go run main.go https://go.dev https://google.com
```

_Boom!_ 💥 O arquivo `output.pdf` aparecerá na sua pasta como se fosse mágica.

#### 2. Modo Servidor (API Power)

Rode sem argumentos para subir o servidor:

```bash
go run main.go
# 🚀 RapidPDF — Web-to-PDF Converter
# 📡 Server listening on :8080
```

Agora você tem superpoderes via HTTP:

- **Gerar PDF**: `POST /generate` com JSON `{"urls": ["..."]}`
- **Documentação**: Acesse [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) e brinque com a API.

### ⚙️ Configuração

Crie um arquivo `.env` para tunar seu RapidPDF:

| Variável          | Descrição                                              | Padrão |
| ----------------- | ------------------------------------------------------ | ------ |
| `MAX_URLS`        | Máximo de URLs permitidas por requisição               | `10`   |
| `TIMEOUT_SECONDS` | Tempo limite (em segundos) para renderizar cada página | `60`   |

### 🛠️ Tecnologias (O Motor)

Debaixo do capô, temos a elite do ecossistema Go:

- 🐹 **Go**: Velocidade e simplicidade.
- 🍸 **Gin**: O framework web supersônico.
- 📜 **Swagger**: Documentação automática.
- 🌐 **Chromedp**: Renderização fiel via Chrome.
- 📄 **pdfcpu**: Cola digital para PDFs.
- 📝 **godotenv**: Gestão de configuração sem dor de cabeça.

---

## 🇺🇸 English

### 🤔 What is this?

Tired of "Save as PDF" one by one? **RapidPDF** is your magical CLI tool that devours URLs, renders them with real-browser precision (thanks, Chrome! 🤖), and stitches them into a single, beautiful PDF. It's like a sticker album of the internet, but actually useful.

### ✨ Superpowers

- **Multitasking**: 1, 10, or 100 URLs? Bring it on.
- **Server Mode**: Run without arguments to launch a powerful REST API. 🛡️
- **Documented**: Swagger UI included, because guessing endpoints is so 2010. 🎩
- **Smart**: Uses the Chrome engine (`chromedp`) to ensure the PDF looks _exactly_ like the website.
- **Organized**: Merges everything into a final `output.pdf` file.
- **Safe**: Validates your URLs so you don't look silly.
- **Configurable**: Adjustable URL limits and Timeout via `.env` (because time is money).

### 🚀 Let's Run It

#### 1. CLI Mode (Classic)

Feed it URLs and watch it fly:

```bash
go run main.go https://go.dev https://google.com
```

_Boom!_ 💥 The `output.pdf` file appears in your folder like magic.

#### 2. Server Mode (API Power)

Run without arguments to launch the server:

```bash
go run main.go
# 🚀 RapidPDF — Web-to-PDF Converter
# 📡 Server listening on :8080
```

Now you have HTTP superpowers:

- **Generate PDF**: `POST /generate` with JSON `{"urls": ["..."]}`
- **Documentation**: Go to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) and play with the API.

### ⚙️ Configuration

Create a `.env` file to tune your RapidPDF:

| Variable          | Description                                  | Default |
| ----------------- | -------------------------------------------- | ------- |
| `MAX_URLS`        | Maximum URLs allowed per request             | `10`    |
| `TIMEOUT_SECONDS` | Timeout (in seconds) for rendering each page | `60`    |

### 🛠️ Tech Stack (The Engine)

Under the hood, we have the elite of the Go ecosystem:

- 🐹 **Go**: Speed and simplicity.
- 🍸 **Gin**: Supersonic web framework.
- 📜 **Swagger**: Automatic documentation.
- 🌐 **Chromedp**: Faithful rendering via Chrome.
- 📄 **pdfcpu**: Digital glue for PDFs.
- 📝 **godotenv**: Headache-free configuration management.

---

### 📜 License

MIT © [Paulo Silva](https://github.com/psilva1982)

_Made with ❤️, code, and maybe too much coffee._
