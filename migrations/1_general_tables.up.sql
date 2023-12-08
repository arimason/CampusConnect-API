-- Tabela de usuário
CREATE TABLE tb_user (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    permission VARCHAR(20) NOT NULL, -- Pode ser 'admin', 'professor' ou 'aluno'
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_user PRIMARY KEY (id)
);

-- Tabela de informações pessoais
CREATE TABLE tb_person (
    id VARCHAR(36)  NOT NULL,
    user_id VARCHAR(36) UNIQUE NOT NULL,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    email VARCHAR(100) UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_person PRIMARY KEY (id),
    CONSTRAINT fk_tb_person_user FOREIGN KEY (user_id) REFERENCES tb_user(id)
);

-- Tabela de eventos
CREATE TABLE tb_event (
    id VARCHAR(36) NOT NULL,
    creator_user_id VARCHAR(36) NOT NULL,
    event_name VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_event PRIMARY KEY (id),
    CONSTRAINT fk_tb_event_user FOREIGN KEY (creator_user_id) REFERENCES tb_user(id)
);

-- Tabela de artigos
CREATE TABLE tb_article (
    id VARCHAR(36) NOT NULL,
    creator_user_id VARCHAR(36) NOT NULL,
    title VARCHAR(100),
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_article PRIMARY KEY (id),
    CONSTRAINT fk_tb_article_user FOREIGN KEY (creator_user_id) REFERENCES tb_user(id)
);

-- Tabela de comentários
CREATE TABLE tb_comment (
    id VARCHAR(36) NOT NULL,
    entity_type VARCHAR(20), -- Pode ser 'article' ou 'event'
    entity_id VARCHAR(36) NOT NULL, -- ID do artigo ou evento vinculado
    user_id VARCHAR(36) NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_comment PRIMARY KEY (id),
    CONSTRAINT fk_tb_comment_user FOREIGN KEY (user_id) REFERENCES tb_user(id),
    CONSTRAINT fk_tb_comment_article FOREIGN KEY (entity_id) REFERENCES tb_article(id) ON DELETE CASCADE,
    CONSTRAINT fk_tb_comment_event FOREIGN KEY (entity_id) REFERENCES tb_event(id) ON DELETE CASCADE
);

-- Outras tabelas necessárias conforme as funcionalidades específicas do seu sistema
