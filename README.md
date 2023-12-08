# CampusConnect-API

## Descrição 

    Layout base: https://github.com/golang-standards/project-layout

### Base URL

``` 

PUBLIC_BASE = /pub
PRIVATE_BASE = /priv

``` 

# ENDPOINTS

### 1-1. [POST]/{PUBLIC_BASE}/user

Endpoint para criar um novo usuário.

#### - _Request_:

O corpo da requisição deve ser um objeto JSON contendo as seguintes informações:

```json
{
"name": "Nome do Usuário",
"email": "usuario@example.com",
"password": "senha_do_usuario"
}
```

#### - _Response_:

O corpo de resposta deve conter:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

### 1-3. [GET]/{PRIVATE_BASE}/user

Endpoint para buscar dados de um usuário.

##### _Request HEADER_:

O HEADER deve conter o token:

```
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

#####  _Response_:

O corpo de resposta em JSON:

```json
{
"name": "Nome do Usuário",
"email": "usuario@example.com",
"password": "senha_do_usuario"
}
```
