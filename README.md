## TODO
- handlers do módulo book
- funcao de sanitizacao de string
- dockerfile
- ci/cd
- testes
    - unitário
    - teste das rotas

## Introdução
O servidor foi desenvolvido conforme as especificações dadas, em Golang e baseado na framework web [Echo](https://github.com/labstack/echo). Além desse, também foram utilizados os seguintes módulos:

- [go-ozzo/ozzo-validation](https://github.com/go-ozzo/ozzo-validation) : para validação de input vinda das requisições
- [po-pg/pg](https://github.com/go-pg/pg): ORM e cliente para utilizar o banco Postgres
- [google/uuid](https://github.com/google/uuid): para trabalhar com identificadores únicos
- [rs/zerolog](https://github.com/rs/zerolog): para lidar com logs do servidor

## Armazenamento
O banco utilizado é um Postgres e ele está hospedado pelo serviço Amazon RDS, na AWS. As credenciais para se conectar com ele estão especificads no arquivo de configuração do app.

## Execução
O servidor pode ser executado a partir do comando `go run .` na pasta raíz do projeto. A intenção é dockerizar e configurar uma pipeline de CICD para uma instância EC2 da AWS que irá executar a aplicação, porém isso ainda não foi feito.

## Testes
TODO