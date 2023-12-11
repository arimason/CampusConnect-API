-- Definindo o enum para permission
CREATE TYPE permission_type AS ENUM ('owner', 'admin', 'teacher','student');

-- Tabela de usuário
CREATE TABLE tb_user (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    permission permission_type, -- Pode ser 'owner', 'admin', 'teacher','student'
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_user PRIMARY KEY (id)
);

-- Definindo o enum para period
CREATE TYPE period_type AS ENUM ('morning', 'evening', 'daytime', 'night', 'all');

-- Tabela de cursos
CREATE TABLE tb_course (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(100) NOT NULL,
    period period_type, -- Pode ser, 'morning', 'evening', 'daytime', 'night', 'all'; uso exclusivo para identificar exatamente o curso do aluno
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_course PRIMARY KEY (id)
);

-- Tabela de informações pessoais
CREATE TABLE tb_person (
    id VARCHAR(36)  NOT NULL,
    user_id VARCHAR(36) UNIQUE NOT NULL,
    course_id VARCHAR(36) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_person PRIMARY KEY (id),
    CONSTRAINT fk_tb_person_user FOREIGN KEY (user_id) REFERENCES tb_user(id),
    CONSTRAINT fk_tb_person_course FOREIGN KEY (course_id) REFERENCES tb_course(id)
);

-- Definindo o enum para period
CREATE TYPE visibility_type AS ENUM ('pub', 'priv');

-- Tabela de eventos
CREATE TABLE tb_event (
    id VARCHAR(36) NOT NULL,
    person_id VARCHAR(36) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    visibility visibility_type, -- Pode ser 'pub', 'priv'
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_event PRIMARY KEY (id),
    CONSTRAINT fk_tb_event_person FOREIGN KEY (person_id) REFERENCES tb_person(id)
);

-- Tabela de artigos
CREATE TABLE tb_article (
    id VARCHAR(36) NOT NULL,
    person_id VARCHAR(36) NOT NULL,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    visibility visibility_type, -- Pode ser 'pub', 'priv'
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_article PRIMARY KEY (id),
    CONSTRAINT fk_tb_article_person FOREIGN KEY (person_id) REFERENCES tb_person(id)
);

-- Defino os valores aceitos para content_type
CREATE TYPE content_type AS ENUM ('article', 'event');

-- Tabela de Relacionamento entre Cursos, Artigos e Eventos
CREATE TABLE tb_course_content (
    course_id VARCHAR(36),
    content_id VARCHAR(36),
    content_type content_type, -- Pode ser 'article' ou 'event'
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP, -- o updated possuirá o mesmo valor que o deleted caso o evento seja excluído
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_course_content PRIMARY KEY (course_id, content_id),
    CONSTRAINT fk_tb_course_content_course FOREIGN KEY (course_id) REFERENCES tb_course(id),
    CONSTRAINT fk_tb_course_content_article FOREIGN KEY (content_id) REFERENCES tb_article(id) DEFERRABLE INITIALLY DEFERRED,
    CONSTRAINT fk_tb_course_content_event FOREIGN KEY (content_id) REFERENCES tb_event(id) DEFERRABLE INITIALLY DEFERRED
);

-- Tabela de inscrição dos alunos nos eventos
CREATE TABLE tb_register_event (
    person_id VARCHAR(36),
    event_id VARCHAR(36),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP, -- o updated possuirá o mesmo valor que o deleted caso o aluno canecele sua inscrição no evento
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_register_event PRIMARY KEY (person_id, event_id),
    CONSTRAINT fk_tb_register_event_person FOREIGN KEY (person_id) REFERENCES tb_person(id),
    CONSTRAINT fk_tb_register_event_event FOREIGN KEY (event_id) REFERENCES tb_event(id)
);

-- Tabela de comentários
CREATE TABLE tb_comment (
    id VARCHAR(36) NOT NULL,
    entity_type content_type, -- Pode ser 'article' ou 'event'
    entity_id VARCHAR(36) NOT NULL, -- ID do artigo ou evento vinculado
    person_id VARCHAR(36) NOT NULL,
    title VARCHAR(100),
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_comment PRIMARY KEY (id),
    CONSTRAINT fk_tb_comment_person FOREIGN KEY (person_id) REFERENCES tb_person(id),
    CONSTRAINT fk_tb_comment_article FOREIGN KEY (entity_id) REFERENCES tb_article(id) ON DELETE CASCADE,
    CONSTRAINT fk_tb_comment_event FOREIGN KEY (entity_id) REFERENCES tb_event(id) ON DELETE CASCADE
);

-- Tabela de documentos que podem ser anexados aos eventos e artigos
CREATE TABLE tb_document (
    id VARCHAR(36) NOT NULL,
    entity_type content_type, -- Pode ser 'article' ou 'event'
    entity_id VARCHAR(36) NOT NULL, -- ID do artigo ou evento vinculado
    document TEXT, -- conteúdo do documento
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT pk_tb_document PRIMARY KEY (id),
    CONSTRAINT fk_tb_document_article FOREIGN KEY (entity_id) REFERENCES tb_article(id) ON DELETE CASCADE,
    CONSTRAINT fk_tb_document_event FOREIGN KEY (entity_id) REFERENCES tb_event(id) ON DELETE CASCADE
);
