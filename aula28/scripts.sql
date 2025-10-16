CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.usuarios
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    login character varying(200) COLLATE pg_catalog."default" NOT NULL,
    senha character varying(200) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT usuarios_pkey PRIMARY KEY (id)
)

INSERT INTO USUARIOS(LOGIN, SENHA) VALUES ('admin','master')
INSERT INTO USUARIOS(LOGIN, SENHA) VALUES ('admin','eb0a191797624dd3a48fa681d3061212')

//gera uuid no postgres
uuid_generate_v4()