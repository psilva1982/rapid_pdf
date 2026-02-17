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

Agora com **Superpoderes de Nuvem**! ☁️ Seus PDFs podem ir direto pro **AWS S3** ou ficar na sua máquina, você manda.

### ✨ Superpoderes

- **Multitarefa**: 1, 10 ou 100 URLs? Manda ver.
- **Modo Servidor**: Rode sem argumentos e ele vira uma API REST poderosa. 🛡️
- **Nuvem ou Terra Firme**: Salva automaticamente no **AWS S3** se configurado, ou na pastinha `./media` se você for _old school_.
- **Documentado**: Swagger UI incluído, porque ninguém merece adivinhar rotas. 🎩
- **Inteligente**: Usa o motor do Chrome (`chromedp`) para garantir que o PDF fique _igualzinho_ ao site.
- **Organizado**: Junta (merge) todas as páginas em um arquivo final.
- **Seguro**: Valida suas URLs para você não passar vergonha.
- **Configurável**: Limites de URLs, Timeout e S3 ajustáveis via `.env`.

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

Se você não configurou o S3, ele vai avisar que está salvando localmente. Sem pânico!

Agora você tem superpoderes via HTTP:

- **Gerar PDF**: `POST /generate` com JSON `{"urls": ["..."]}`
- **Resposta**: Ele te devolve uma URL bonitinha, seja do S3 ou local! \o/

  ```json
  {
    "url": "https://meu-bucket.s3.us-east-1.amazonaws.com/pdfs/2023/10/arquivo.pdf"
  }
  ```

  _(Ou `/media/arquivo.pdf` se estiver rodando local)_

- **Documentação**: Acesse [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) e brinque com a API.

#### 3. Modo Docker 🐳

A maneira mais fácil de rodar sem instalar dependências!

**Puxar e rodar do DockerHub:**

```bash
docker pull severos/rapid-pdf:latest
docker run -d -p 8080:8080 -v $(pwd)/media:/app/media severos/rapid-pdf:latest
```

**Ou usar Docker Compose:**

```bash
docker-compose up -d
```

Pronto! Acesse `http://localhost:8080/swagger/index.html` para usar a API.

📚 **Documentação completa**: Veja [DOCKER.md](DOCKER.md) para configurações avançadas, deployment em produção e troubleshooting.

### ⚙️ Configuração

Crie um arquivo `.env` para tunar seu RapidPDF. Agora com chaves do S3 e Porta customizável!

| Variável                 | Descrição                                                 | Padrão    |
| :----------------------- | :-------------------------------------------------------- | :-------- |
| `PORT`                   | Porta do servidor web superônico                          | `8080`    |
| `MAX_URLS`               | Máximo de URLs permitidas por requisição                  | `10`      |
| `TIMEOUT_SECONDS`        | Tempo limite (em segundos) para renderizar cada página    | `60`      |
| `PAGE_LOAD_WAIT_SECONDS` | Tempo de espera (em segundos) após carregamento da página | `5`       |
| `AWS_S3_BUCKET`          | Nome do seu balde (bucket) no S3 🪣                       | _(vazio)_ |
| `AWS_S3_REGION`          | Região da AWS (ex: `us-east-1`)                           | _(vazio)_ |
| `AWS_S3_ACCESS_KEY`      | Sua chave de acesso AWS 🔑                                | _(vazio)_ |
| `AWS_S3_SECRET_KEY`      | Seu segredo AWS 🤫                                        | _(vazio)_ |

> **Dica de Mestre**: Se não preencher as variáveis da AWS, o RapidPDF assume o modo "Hacker de Garagem" e salva tudo na pasta `./media`.

### 🛠️ Tecnologias (O Motor)

Debaixo do capô, temos a elite do ecossistema Go:

- 🐹 **Go**: Velocidade e simplicidade.
- 🍸 **Gin**: O framework web supersônico.
- ☁️ **AWS SDK v2**: Falando a língua das nuvens.
- 📜 **Swagger**: Documentação automática.
- 🌐 **Chromedp**: Renderização fiel via Chrome.
- 📄 **pdfcpu**: Cola digital para PDFs.
- 📝 **godotenv**: Gestão de configuração sem dor de cabeça.

---

## 🇺🇸 English

### 🤔 What is this?

Tired of "Save as PDF" one by one? **RapidPDF** is your magical CLI tool that devours URLs, renders them with real-browser precision (thanks, Chrome! 🤖), and stitches them into a single, beautiful PDF. It's like a sticker album of the internet, but actually useful.

Now with **Cloud Superpowers**! ☁️ Your PDFs can go straight to **AWS S3** or stay on your machine. You're the boss.

### ✨ Superpowers

- **Multitasking**: 1, 10, or 100 URLs? Bring it on.
- **Server Mode**: Run without arguments to launch a powerful REST API. 🛡️
- **Cloud or Ground**: Automatically saves to **AWS S3** if configured, or to `./media` if you're keeping it old school.
- **Documented**: Swagger UI included, because guessing endpoints is so 2010. 🎩
- **Smart**: Uses the Chrome engine (`chromedp`) to ensure the PDF looks _exactly_ like the website.
- **Organized**: Merges everything into a final file.
- **Safe**: Validates your URLs so you don't look silly.
- **Configurable**: Adjustable URL limits, Timeout, and S3 settings via `.env`.

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

If S3 isn't set up, it'll tell you it's saving locally. Don't panic!

Now you have HTTP superpowers:

- **Generate PDF**: `POST /generate` with JSON `{"urls": ["..."]}`
- **Response**: It hands you back a shiny URL, either from S3 or local! \o/

  ```json
  {
    "url": "https://my-bucket.s3.us-east-1.amazonaws.com/pdfs/2023/10/file.pdf"
  }
  ```

  _(Or `/media/file.pdf` if running locally)_

- **Documentation**: Go to [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) and play with the API.

#### 3. Docker Mode 🐳

The easiest way to run without installing dependencies!

**Pull and run from DockerHub:**

```bash
docker pull severos/rapid-pdf:latest
docker run -d -p 8080:8080 -v $(pwd)/media:/app/media severos/rapid-pdf:latest
```

**Or use Docker Compose:**

```bash
docker-compose up -d
```

Done! Go to `http://localhost:8080/swagger/index.html` to use the API.

📚 **Full Documentation**: See [DOCKER.md](DOCKER.md) for advanced configuration, production deployment, and troubleshooting.

### ⚙️ Configuration

Create a `.env` file to tune your RapidPDF. Now with S3 keys and custom Port!

| Variable                 | Description                                  | Default   |
| :----------------------- | :------------------------------------------- | :-------- |
| `PORT`                   | Server port (where the magic happens)        | `8080`    |
| `MAX_URLS`               | Maximum URLs allowed per request             | `10`      |
| `TIMEOUT_SECONDS`        | Timeout (in seconds) for rendering each page | `60`      |
| `PAGE_LOAD_WAIT_SECONDS` | Wait time (in seconds) after page load       | `5`       |
| `AWS_S3_BUCKET`          | Your S3 bucket name 🪣                       | _(empty)_ |
| `AWS_S3_REGION`          | AWS Region (e.g., `us-east-1`)               | _(empty)_ |
| `AWS_S3_ACCESS_KEY`      | Your AWS Access Key 🔑                       | _(empty)_ |
| `AWS_S3_SECRET_KEY`      | Your AWS Secret Key 🤫                       | _(empty)_ |

> **Pro Tip**: If you leave the AWS variables empty, RapidPDF goes into "Garage Hacker" mode and saves everything to the `./media` folder.

### 🛠️ Tech Stack (The Engine)

Under the hood, we have the elite of the Go ecosystem:

- 🐹 **Go**: Speed and simplicity.
- 🍸 **Gin**: Supersonic web framework.
- ☁️ **AWS SDK v2**: Speaking the language of the clouds.
- 📜 **Swagger**: Automatic documentation.
- 🌐 **Chromedp**: Faithful rendering via Chrome.
- 📄 **pdfcpu**: Digital glue for PDFs.
- 📝 **godotenv**: Headache-free configuration management.

---

### 📜 License

MIT © [Paulo Silva](https://github.com/psilva1982)

_Made with ❤️, code, and maybe too much coffee._
