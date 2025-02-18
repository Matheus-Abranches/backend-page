# 📌 backend-page

## 🚀 Sobre o Projeto
Esta é uma API REST desenvolvida em **Go** que implementa um CRUD (Create, Read, Update, Delete) para gerenciar usuários e e-mails enviados. O projeto utiliza **Gin** como framework web e um banco de dados relacional gerenciado com **SQL**.

## 🛠️ Tecnologias Utilizadas
- **Go** - Linguagem de programação principal do projeto
- **Gin** - Framework para construção de APIs
- **SQL** - Banco de dados relacional
- **DBeaver** - Ferramenta de gerenciamento do banco
- **Docker** - Para conteinerização e facilitação do ambiente

## 🏗️ Fluxo de Trabalho (Workflow)

```mermaid
flowchart LR
    A[Cliente / Frontend] -->|Requisição HTTP (CRUD)| B(Controller)
    B --> C(UseCase<br>Regras de Negócio)
    C --> D(Repository<br>Operações CRUD)
    D --> E[(Database<br>Usuário, SentEmail, etc.)]
    E --> D
    D --> C
    C --> B
    B -->|Resposta HTTP| A

## ⚙️ Como Rodar o Projeto

### 📌 1. Clonar o Repositório
```sh
git clone https://github.com/seuusuario/seurepositorio.git
cd seurepositorio
```

### 🐳 2. Subir o Ambiente com Docker
```sh
docker-compose up --build
```

### ▶️ 3. Rodar a API Manualmente
Se não for usar Docker, execute:
```sh
go run main.go
```

## 📌 Rotas da API

### Usuários
- **GET /users** - Buscar todos os usuários
- **POST /users** - Criar um novo usuário
- **PUT /users/:id** - Atualizar um usuário
- **DELETE /users/:id** - Remover um usuário (soft delete)

### E-mails Enviados
- **GET /sentemails** - Buscar todos os e-mails enviados
- **POST /sentemails** - Registrar um novo e-mail enviado
- **PUT /sentemails/:id** - Atualizar um e-mail enviado
- **DELETE /sentemails/:id** - Remover um e-mail enviado

## 📞 Contato
Caso tenha alguma dúvida, sugestão ou precise de suporte, entre em contato:

📧 **E-mail:** matheusabrancheslimatav@gmail.com  
💼 **LinkedIn:** [Matheus Abranches Lima Tavares](https://www.linkedin.com/in/matheusabranches/)
📱 **Telefone/Celular:** [+55 (32) 98845-1328]

Estou à disposição para ajudar! 🚀

---
👨‍💻 **Desenvolvido por [Matheus Abranches Lima Tavares](https://github.com/Matheus-Abranches)**

