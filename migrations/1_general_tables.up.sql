CREATE TABLE tb_user (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(255),
    Email VARCHAR(255),
    Password VARCHAR(255),
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    DeletedAt TIMESTAMP
);

-- CREATE TABLE IF NOT EXISTS tb_sys_job(
-- 	id int not null,
--   kind varchar(20) not null,
-- 	status varchar(20) not null,
-- 	processed bool not null,
--   configuration jsonb not null,
--   with_error bool not null,
-- 	result jsonb,
-- 	created_at timestamptz not null default now(),
-- 	updated_at timestamptz,
-- 	deleted_at timestamptz,
-- 	constraint pk_tb_sys_job primary key (id)
-- );