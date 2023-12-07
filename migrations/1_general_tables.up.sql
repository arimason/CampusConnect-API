CREATE TABLE tb_user (
    id varchar(36) not null,
    name VARCHAR(255) not null,
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    created_at TIMESTAMP not null DEFAULT now(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    constraint pk_tb_user primary key (id)
);