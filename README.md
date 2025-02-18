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

1. **Cliente/Frontend**  
   O usuário (via navegador ou ferramenta como Postman) faz uma requisição HTTP (CRUD) para criar, ler, atualizar ou deletar dados.

2. **Controller**  
   Recebe a requisição HTTP, faz a validação inicial dos dados e os repassa para o caso de uso adequado.

3. **UseCase (Regras de Negócio)**  
   Contém a lógica de negócio. Aqui são realizadas validações mais complexas, regras específicas e orquestração do fluxo. Em seguida, o UseCase solicita ao repositório (Repository) que interaja com o banco de dados.

4. **Repository**  
   Responsável pela comunicação com o banco de dados. Executa as operações CRUD (Create, Read, Update, Delete) nas tabelas correspondentes às entidades (Usuário, SentEmail etc.).

5. **Database**  
   É onde os dados de Usuários e E-mails Enviados (SentEmail) são efetivamente armazenados. Após as operações, o repositório retorna o resultado ao UseCase.

6. **Retorno da Resposta**  
   O UseCase retorna o resultado final (sucesso ou erro) para o Controller, que então envia a resposta HTTP ao Cliente/Frontend.


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

