# CampusConnect-API

## Layout base: https://github.com/golang-standards/project-layout


### ENDPOINTS

#### Criar Usuário

Endpoint para criar um novo usuário.

- **Método:** POST
- **Path:** `/pub/user`

##### Request Body

O corpo da requisição deve ser um objeto JSON contendo as seguintes informações:

```json
{
"name": "Nome do Usuário",
"email": "usuario@example.com",
"password": "senha_do_usuario"
}
```

##### Request Body

O corpo de resposta deve conter:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```
