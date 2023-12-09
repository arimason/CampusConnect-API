-- Tabela de documentos que podem ser anexados aos eventos e artigos
DROP TABLE IF EXISTS tb_document;

-- Tabela de comentários
DROP TABLE IF EXISTS tb_comment;

-- Tabela de Relacionamento entre Cursos, Artigos e Eventos
DROP TABLE IF EXISTS tb_course_content;

-- Defino os valores aceitos para content_type
DROP TYPE IF EXISTS content_type;

-- Tabela de artigos
DROP TABLE IF EXISTS tb_article;

-- Tabela de eventos
DROP TABLE IF EXISTS tb_event;

-- Definindo o enum para visibility
DROP TYPE IF EXISTS visibility_type;

-- Tabela de informações pessoais
DROP TABLE IF EXISTS tb_person;

-- Tabela de cursos
DROP TABLE IF EXISTS tb_course;

-- Definindo o enum para period
DROP TYPE IF EXISTS period_type;

-- Tabela de usuário
DROP TABLE IF EXISTS tb_user;

-- Definindo o enum para permission
DROP TYPE IF EXISTS permission_type;
