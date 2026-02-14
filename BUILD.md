# 🏗️ Build & Run Masterclass

> **Where code becomes prowess.** / **Onde o código vira poder.**

[![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=flat-square&logo=go)](https://go.dev/)

---

### 🌐 Select Language / Selecione o Idioma

- [🇧🇷 Português Brasileiro (PT-BR)](#-português-brasileiro)
- [🇺🇸 English (EN)](#-english)

---

## 🇧🇷 Português Brasileiro

### 📋 O Kit de Sobrevivência (Pré-requisitos)

Antes de começar a brincadeira, você precisa ter:

- **Go**: Versão 1.25 ou superior. (Afinal, queremos o que há de mais moderno!)
  - [Baixar Go](https://go.dev/dl/)
- **Google Chrome** (ou Chromium): O motor da nossa Ferrari.
  - O app tenta achar sozinho, igual mágica.
- **Git**: Para baixar o código sem sofrer.

### 🛠️ Instalação Relâmpago

1.  **Clone o repo**:

    ```bash
    git clone https://github.com/psilva1982/rapid_pdf.git
    cd rapid_pdf
    ```

2.  **Chame os reforços (Dependências)**:
    ```bash
    go mod download
    ```

### 🏗️ Construindo o Monstro (Build)

Quer criar um executável para levar pra casa?

**Passo 0: Gere a documentação (Swagger)**
Se você mexeu na API, rode isso antes ou o compilador vai chorar:

```bash
swag init
```

**Passo 1: Compile**

**Linux / macOS:**

```bash
go build -o rapid_pdf main.go
```

**Windows:**

```powershell
go build -o rapid_pdf.exe main.go
```

Pronto! Agora você tem um binário fresquinho na pasta.

### 🚀 Decolando (Execução)

**Modo "API Server" (Padrão):**

```bash
# Roda na porta 8080
./rapid_pdf
```

**Modo "CLI One-Shot" (Clássico):**

```bash
# Converte e sai
./rapid_pdf https://exemplo.com.br
```

### ⚙️ Ajustes Finos (Configuração)

Crie um arquivo `.env` e mande ver nas configs:

| Variável   | O que faz?                                       | Padrão |
| :--------- | :----------------------------------------------- | :----- |
| `MAX_URLS` | Quantos sites você aguenta converter de uma vez? | `10`   |
| `TIMEOUT`  | Tempo (em seg) antes de desistir se a net cair.  | `30`   |

**Exemplo `.env`**:

```env
MAX_URLS=42
TIMEOUT=60
```

### 🧪 Testando Tudo

Para ter certeza que nada quebrou:

```bash
go test -v ./...
```

### ❌ Deu Ruim? (Troubleshooting)

| Problema                    | Possível Causa    | Solução                               |
| :-------------------------- | :---------------- | :------------------------------------ |
| `executable file not found` | Cadê o Chrome?    | Instala o Chrome aí, chefia.          |
| `context deadline exceeded` | Internet discada? | Aumenta o `TIMEOUT` ou checa o Wi-Fi. |
| `too many URLs`             | Calma, jovem!     | Aumenta o `MAX_URLS` no `.env`.       |
| `docs package not found`    | Esqueceu o swag?  | Roda `swag init` antes do build!      |

---

## 🇺🇸 English

### 📋 Survival Kit (Prerequisites)

Before joining the party, make sure you have:

- **Go**: Version 1.25 or higher. (We want the cutting edge!)
  - [Download Go](https://go.dev/dl/)
- **Google Chrome** (or Chromium): The engine of our Ferrari.
  - The app tries to find it automagically.
- **Swag CLI**: For generating docs. `go install github.com/swaggo/swag/cmd/swag@latest`.
- **Git**: To get the code painlessly.

### 🛠️ Flash Installation

1.  **Clone the repo**:

    ```bash
    git clone https://github.com/psilva1982/rapid_pdf.git
    cd rapid_pdf
    ```

2.  **Call for backup (Dependencies)**:
    ```bash
    go mod download
    ```

### 🏗️ Building the Beast

Want a standalone executable to take home?

**Step 0: Generate Docs**
Don't skip this or the build will fail:

```bash
swag init
```

**Step 1: Compile**

**Linux / macOS:**

```bash
go build -o rapid_pdf main.go
```

**Windows:**

```powershell
go build -o rapid_pdf.exe main.go
```

Done! You now have a fresh binary in your folder.

### 🚀 Liftoff (Running)

**"API Server" Mode (Default):**

```bash
# Listens on :8080
./rapid_pdf
```

**"CLI One-Shot" Mode (Classic):**

```bash
# Converts and exits
./rapid_pdf https://example.com
```

### ⚙️ Fine Tuning (Configuration)

Create a `.env` file and tweak the settings:

| Variable   | What does it do?                                 | Default |
| :--------- | :----------------------------------------------- | :------ |
| `MAX_URLS` | How many sites can you handle at once?           | `10`    |
| `TIMEOUT`  | Time (in sec) before giving up if the net fails. | `30`    |

**Example `.env`**:

```env
MAX_URLS=42
TIMEOUT=60
```

### 🧪 Testing Everything

To make sure nothing blew up:

```bash
go test -v ./...
```

### ❌ Oops? (Troubleshooting)

| Issue                       | Possible Cause      | Solution                           |
| :-------------------------- | :------------------ | :--------------------------------- |
| `executable file not found` | Where is Chrome?    | Install Chrome, boss.              |
| `context deadline exceeded` | Dial-up internet?   | Increase `TIMEOUT` or check Wi-Fi. |
| `too many URLs`             | Easy there, cowboy! | Increase `MAX_URLS` in `.env`.     |

---

> **Note**: This app uses `chromedp` (headless Chrome). If running in Docker, you might need special flags (or just a miracle).
