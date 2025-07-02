# ✉️ emailn

**emailn** é um projeto de aprendizado desenvolvido em Go (Golang) com o objetivo de explorar conceitos de backend como APIs REST, envio de e-mails e SMS, e manipulação de dados em campanhas de marketing.

Este sistema permite cadastrar **contatos**, **campanhas** e realizar o **envio de mensagens** aos contatos associados a cada campanha.

> ⚠️ **Projeto em desenvolvimento**: Este repositório está sendo desenvolvido com fins educacionais. Funcionalidades podem estar incompletas ou instáveis.

## 🎯 Funcionalidades

- 👤 Cadastro de contatos com nome, e-mail e telefone
- 📣 Criação de campanhas com título e conteúdo da mensagem
- 🔗 Associação de contatos a campanhas
- ✉️ Envio de e-mails para os contatos de uma campanha
- 📱 Envio de SMS (com suporte planejado ou simulado)
- 📦 API RESTful para gerenciamento dos dados

## 🚀 Tecnologias utilizadas

- [Golang](https://golang.org/) como linguagem principal
- [Gin](https://github.com/gin-gonic/gin) para criação da API
- [SMTP](https://pkg.go.dev/net/smtp) para envio de e-mails
- Integração futura com serviços SMS como Twilio

## 🛠 Como executar localmente

> Pré-requisitos:
> - Go instalado (versão 1.20 ou superior)
> - Banco de dados PostgreSQL

1. Clone o projeto:
   ```bash
   git clone https://github.com/seu-usuario/emailn.git
   cd emailn
