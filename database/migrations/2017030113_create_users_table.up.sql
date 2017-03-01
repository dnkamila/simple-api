CREATE TABLE "users" (
    id bigint NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(50) NOT NULL
);

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY "users" ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

SELECT pg_catalog.setval('users_id_seq', 1, false);

ALTER TABLE ONLY "users" ADD CONSTRAINT user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY "users" ADD CONSTRAINT user_username_key UNIQUE (username);