# Descrição do Banco de Dados

## Tabela `tb_user`

Armazena informações sobre os usuários do sistema.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do usuário                      |
| name          | VARCHAR(50)      | Nome do usuário                                     |
| email         | VARCHAR(100)     | Endereço de e-mail único do usuário                 |
| password      | VARCHAR(100)     | Senha do usuário                                    |
| permission    | permission_type  | Tipo de permissão do usuário ('owner', 'admin', 'teacher', 'student') |
| created_at    | TIMESTAMP         | Data e hora de criação do usuário                   |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

---

## Tabela `tb_course`

Armazena informações sobre os cursos disponíveis.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do curso                        |
| name          | VARCHAR(100)     | Nome do curso                                       |
| period        | period_type      | Período do curso ('morning', 'evening', 'daytime', 'night', 'all') |
| created_at    | TIMESTAMP         | Data e hora de criação do curso                     |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (m‌arcada)                   |

---

## Tabela `tb_person`

Armazena informações pessoais vinculadas aos usuários.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único das informações pessoais        |
| user_id       | VARCHAR(36)      | Identificador único do usuário vinculado            |
| course_id     | VARCHAR(36)      | Identificador único do curso vinculado              |
| first_name    | VARCHAR(50)      | Primeiro nome do usuário                            |
| last_name     | VARCHAR(50)      | Sobrenome do usuário                                |
| created_at    | TIMESTAMP         | Data e hora de criação das informações pessoais     |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

*Restrições de chave estrangeira:*
- `fk_tb_person_user`: Referencia `tb_user` através de `user_id`.
- `fk_tb_person_course`: Referencia `tb_course` através de `course_id`.

---

## Tabela `tb_event`

Armazena informações sobre eventos criados pelos usuários.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do evento                       |
| user_id       | VARCHAR(36)      | Identificador único do usuário criador do evento     |
| title         | VARCHAR(100)     | Título do evento                                    |
| content       | TEXT             | Conteúdo do evento                                  |
| visibility    | visibility_type  | Visibilidade do evento ('pub', 'priv')              |
| created_at    | TIMESTAMP         | Data e hora de criação do evento                    |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

*Restrição de chave estrangeira:*
- `fk_tb_event_user`: Referencia `tb_user` através de `user_id`.

---

## Tabela `tb_article`

Armazena informações sobre artigos criados pelos usuários.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do artigo                       |
| user_id       | VARCHAR(36)      | Identificador único do usuário criador do artigo     |
| title         | VARCHAR(100)     | Título do artigo                                    |
| content       | TEXT             | Conteúdo do artigo                                  |
| created_at    | TIMESTAMP         | Data e hora de criação do artigo                    |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

*Restrição de chave estrangeira:*
- `fk_tb_article_user`: Referencia `tb_user` através de `user_id`.

---

## Tabela `tb_course_content`

Estabelece relações entre cursos, artigos e eventos.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| course_id     | VARCHAR(36)      | Identificador único do curso vinculado              |
| content_id    | VARCHAR(36)      | Identificador único do artigo ou evento vinculado   |
| content_type  | content_type      | Tipo de conteúdo ('article' ou 'event')             |

*Restrições de chave estrangeira:*
- `fk_tb_course_content_course`: Referencia `tb_course` através de `course_id`.
- `fk_tb_course_content_article`: Referencia `tb_article` através de `content_id` com verificação adiada.
- `fk_tb_course_content_event`: Referencia `tb_event` através de `content_id` com verificação adiada.

---

## Tabela `tb_comment`

Armazena comentários feitos pelos usuários em artigos e eventos.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do comentário                   |
| entity_type   | content_type      | Tipo de entidade alvo do comentário ('article' ou 'event') |
| entity_id     | VARCHAR(36)      | Identificador único do artigo ou evento vinculado   |
| user_id       | VARCHAR(36)      | Identificador único do usuário que fez o comentário  |
| title         | VARCHAR(100)     | Título do comentário (opcional)                     |
| content       | TEXT             | Conteúdo do comentário                              |
| created_at    | TIMESTAMP         | Data e hora de criação do comentário                |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

*Restrição de chave estrangeira:*
- `fk_tb_comment_user`: Referencia `tb_user` através de `user_id`.
- `fk_tb_comment_article`: Referencia `tb_article` através de `entity_id` com exclusão em cascata.
- `fk_tb_comment_event`: Referencia `tb_event` através de `entity_id` com exclusão em cascata.

---

## Tabela `tb_document`

Armazena documentos anexados a artigos e eventos pelos usuários.

| Coluna        | Tipo             | Descrição                                           |
| ------------- | ---------------- | --------------------------------------------------- |
| id            | VARCHAR(36)      | Identificador único do documento                   |
| entity_type   | content_type      | Tipo de entidade alvo do documento ('article' ou 'event') |
| entity_id     | VARCHAR(36)      | Identificador único do artigo ou evento vinculado   |
| user_id       | VARCHAR(36)      | Identificador único do usuário que anexou o documento |
| document      | TEXT             | Conteúdo do documento                               |
| created_at    | TIMESTAMP         | Data e hora de criação do documento                |
| updated_at    | TIMESTAMP         | Data e hora da última atualização                   |
| deleted_at    | TIMESTAMP         | Data e hora da exclusão (marcada)                   |

*Restrição de chave estrangeira:*
- `fk_tb_document_user`: Referencia `tb_user` através de `user_id`.
- `fk_tb_document_article`: Referencia `tb_article` através de `entity_id` com exclusão em cascata.
- `fk_tb_document_event`: Referencia `tb_event` através de `entity_id` com exclusão em cascata.

---

Este é um esquema básico do banco de dados, e podem ser necessários ajustes conforme a evolução do sistema e requisitos específicos. Certifique-se de adaptar o esquema conforme necessário.
