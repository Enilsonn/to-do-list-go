# 🧩 Go To-Do List API

Uma API RESTful simples e escalável para gerenciamento de **tasks**,
construída em **Go (Golang)** com arquitetura limpa, repositórios, DI e
roteamento eficiente utilizando **Chi**.

------------------------------------------------------------------------

## 📛 Badges

![Go Version](https://img.shields.io/badge/Go-1.21-blue)
![Status](https://img.shields.io/badge/Build-Passing-brightgreen)
![License](https://img.shields.io/badge/License-MIT-green) ![Made
With](https://img.shields.io/badge/Made%20With-Go-blue?logo=go)
![Contributions
Welcome](https://img.shields.io/badge/Contributions-Welcome-orange)

------------------------------------------------------------------------

## 📚 Tabela de Conteúdos

-   [✨ Funcionalidades](#-funcionalidades)
-   [🛠️ Stack de Tecnologias](#️-stack-de-tecnologias)
-   [🏛️ Arquitetura do Projeto](#-arquitetura-do-projeto)
-   [🧱 Diagrama da Arquitetura](#-diagrama-da-arquitetura)
-   [🚀 Como Executar](#-como-executar)
-   [🔌 API Endpoints](#-api-endpoints)
-   [🎯 Exemplo de Requisições](#-exemplo-de-requisições)
-   [🧭 Roadmap](#-roadmap)
-   [🤝 Como Contribuir](#-como-contribuir)
-   [📄 Licença](#-licença)

------------------------------------------------------------------------

## ✨ Funcionalidades

-   Criar tasks
-   Listar todas as tasks
-   Buscar task por ID
-   Atualizar task (incluindo status "done")
-   Deletar task

------------------------------------------------------------------------

## 🛠️ Stack de Tecnologias

-   **Go (Golang)**
-   **Chi Router v5**
-   **SQLite** (*serverless DB*)
-   **database/sql**
-   **context.Context**

------------------------------------------------------------------------

## 🏛️ Arquitetura do Projeto

    /
    ├── go.mod
    ├── main.go
    ├── tasks.db
    └── internal/
        ├── handlers/
        │   └── handlers.go
        ├── models/
        │   └── task.go
        ├── repository/
        │   ├── interface.go
        │   ├── repository.go
        │   └── table-task.go
        └── utils/
            └── respondJSON.go

------------------------------------------------------------------------

## 🧱 Diagrama da Arquitetura

``` mermaid
flowchart TD
    A[Cliente HTTP] --> B[Router Chi]
    B --> C[Handler]
    C --> D[Repository Interface]
    D --> E[Repository SQLite]
    E --> F[(SQLite DB)]
```

------------------------------------------------------------------------

## 🚀 Como Executar

### 1️⃣ Pré-requisitos

-   Go **1.18+**

------------------------------------------------------------------------

### 2️⃣ Clonar o Repositório

``` bash
git clone https://github.com/SEU-USUARIO/SEU-REPOSITORIO.git
cd SEU-REPOSITORIO
```

------------------------------------------------------------------------

### 3️⃣ Instalar Dependências

``` bash
go mod tidy
```

------------------------------------------------------------------------

### 4️⃣ Executar o Servidor

``` bash
go run .
```

Servidor disponível em:\
👉 **http://localhost:8080**

------------------------------------------------------------------------

## 🔌 API Endpoints

Base URL:

    http://localhost:8080/tasks

  ---------------------------------------------------------------------------------------------------------
  Método    Endpoint         Descrição                   Exemplo JSON
  --------- ---------------- --------------------------- --------------------------------------------------
  POST      /tasks           Criar task                  `{ "title": "Estudar Go", "done": false }`

  GET       /tasks           Listar tasks                ---

  GET       /tasks/{id}      Buscar por ID               ---

  PUT       /tasks/{id}      Atualizar task              `{ "title": "Estudar Go (Done)", "done": true }`

  DELETE    /tasks/{id}      Deletar task                ---
  ---------------------------------------------------------------------------------------------------------

------------------------------------------------------------------------

## 🎯 Exemplo de Requisições

### ➕ Criar Task

``` bash
curl -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d '{"title": "Ler documentação do Go", "done": false}'
```

### 📋 Listar Tasks

``` bash
curl http://localhost:8080/tasks
```

### 🔄 Atualizar Task

``` bash
curl -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d '{"title": "Ler documentação do Go (done)", "done": true}'
```

### ❌ Deletar Task

``` bash
curl -X DELETE http://localhost:8080/tasks/1
```

------------------------------------------------------------------------

## 🧭 Roadmap

-   [ ] Criar testes automatizados\
-   [ ] Adicionar Swagger / OpenAPI\
-   [ ] Dockerizar o projeto\
-   [ ] Criar autenticação com JWT\
-   [ ] Adicionar GitHub Actions

------------------------------------------------------------------------

## 🤝 Como Contribuir

1.  Faça um **fork**
2.  Crie uma branch:\
    `git checkout -b minha-feature`
3.  Commit suas mudanças\
4.  Abra um **Pull Request**

Contribuições são sempre bem-vindas! 💙

------------------------------------------------------------------------

## 📄 Licença

Distribuído sob a licença **MIT**.\
Sinta-se livre para usar, modificar e contribuir.

------------------------------------------------------------------------

Desenvolvido com 💙 em **Go**
