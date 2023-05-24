CREATE TABLE IF NOT EXISTS urls
(
	id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
	long text,
	short character varying(10),
	CONSTRAINT urls_pkey PRIMARY KEY (id)
);