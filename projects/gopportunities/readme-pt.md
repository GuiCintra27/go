# GoJob Opportunities API 🇧🇷

<p align="center">
  <img src="./assets/GopportunitiesHeader.svg" alt="GoJob Header">
</p>

Este projeto é uma API moderna de oportunidades de emprego construída usando Golang. A API é desenvolvida com Go-Gin como roteador, GoORM para comunicação de banco de dados, SQLite como banco de dados e Swagger para documentação e teste de API. O projeto segue uma estrutura de pacotes moderna para manter a base de código organizada e de fácil manutenção.

---

## Características

- Introdução ao Golang e construção de APIs modernas
- Configuração do ambiente de desenvolvimento para criação da API
- Usando Go-Gin como roteador para gerenciamento de rotas
- Implementando SQLite como banco de dados para a API
- Utilização do GoORM para comunicação com o banco de dados
- Integração do Swagger para documentação e testes de API
- Criação de uma estrutura de pacotes moderna para organização do projeto
- Implementação de uma API completa de oportunidades de emprego com endpoints para busca, criação, edição e exclusão de oportunidades
- Implementação de testes automatizados para garantir a qualidade da API

## Versão Disponível

Você pode verificar a documentação e realizar testes na API visitando a versão live hospedada em [gopportunities.excalidevs.com](https://gopportunities.excalidevs.com/swagger/index.html).

## Variáveis de ambiente

Este projeto usa uma variável de ambiente para ser executado, verifique o arquivo .env.example para ver o exemplo.

## Comandos Makefile

O projeto inclui um Makefile para ajudá-lo a gerenciar tarefas comuns com mais facilidade. Aqui está uma lista dos comandos disponíveis e uma breve descrição do que eles fazem:

- `make run`: Roda a aplicação sem gerar a documentação da API.
- `make run-with-docs`: Gera a documentação da API, usando Swagger, e então roda a aplicação.
- `make build`: Gera a versão de produção e cria o arquivo executável, nomeado `gopportunities`.
- `make docs`: Gera apenas a documentação da API, utilizando o Swagger.
- `make clean`: Remove o executável `gopportunities`, e deleta o diretório `./docs`.

Para usar esses comandos, basta digitar `make` seguido do comando desejado em seu terminal. Por exemplo, para executar o projeto:

```sh
make run
```

## Ferramentas usadas

Este projeto utiliza as seguintes ferramentas:

- [Golang](https://golang.org/) para desenvolvimento back-end
- [Go-Gin](https://github.com/gin-gonic/gin) para gerenciamento de rotas
- [GoORM](https://gorm.io/) para comunicação com banco de dados
- [SQLite](https://www.sqlite.org/index.html) como banco de dados
- [Swagger](https://swagger.io/) para documentação e testes de API

## Como Acessar

Depois que a API estiver em execução, você poderá usar a UI do Swagger para interagir com os endpoints para pesquisar, criar, editar e excluir oportunidades de emprego. A API pode ser acessada em `http://localhost:$PORT/swagger/index.html`.

Padrão $PORT se não for fornecido=8080.

---

## Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo LICENSE.md para obter detalhes.