CREATE TABLE "user" (
    id bigint NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(50) NOT NULL
);

CREATE SEQUENCE user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE ONLY "user" ALTER COLUMN id SET DEFAULT nextval('user_id_seq'::regclass);

SELECT pg_catalog.setval('user_id_seq', 1, false);

ALTER TABLE ONLY "user" ADD CONSTRAINT user_pkey PRIMARY KEY (id);

ALTER TABLE ONLY "user" ADD CONSTRAINT user_username_key UNIQUE (username);