# 📝 To Do List API (em Go)

Este é um projeto pessoal de uma API RESTful para gerenciamento de tarefas (To-Do List), desenvolvida com **Go (Golang)**. O objetivo do projeto é praticar boas práticas de organização de projetos Go, incluindo uso de `Gin`, `GORM`, `SQLite`, geração automática de documentação com **Swagger**, e automação com **Makefile**.

---

## 🚀 Tecnologias Utilizadas

- [Go] – Linguagem de programação
- [Gin] – Framework web
- [GORM] – ORM para Go
- [SQLite] – Banco de dados
- [Swag] – Geração de documentação Swagger
- Makefile para automação de tarefas

---

## 📁 Estrutura básica

```bash
.
├── internal/          # Código da aplicação (db, handlers, schemas...)
├── docs/              # Documentação Swagger (gerada)
├── Makefile           # Comandos automatizados
├── go.mod / go.sum    # Dependências
└── README.md