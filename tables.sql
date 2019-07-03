--
-- PostgreSQL database dump
--

-- Dumped from database version 9.4.23
-- Dumped by pg_dump version 11.4 (Ubuntu 11.4-0ubuntu0.19.04.1)

-- Started on 2019-07-03 10:50:53 -03

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 2 (class 3079 OID 16428)
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- TOC entry 2068 (class 0 OID 0)
-- Dependencies: 2
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 177 (class 1259 OID 16417)
-- Name: avatar; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.avatar (
    id integer NOT NULL,
    url character varying(255) NOT NULL,
    type character varying(20) NOT NULL
);


ALTER TABLE public.avatar OWNER TO postgres;

--
-- TOC entry 176 (class 1259 OID 16415)
-- Name: avatar_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.avatar_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.avatar_id_seq OWNER TO postgres;

--
-- TOC entry 2069 (class 0 OID 0)
-- Dependencies: 176
-- Name: avatar_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.avatar_id_seq OWNED BY public.avatar.id;


--
-- TOC entry 175 (class 1259 OID 16406)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(255) NOT NULL,
    cpf character varying(15) NOT NULL,
    avatar integer NOT NULL,
    datastart character varying NOT NULL,
    uuiduser uuid DEFAULT public.uuid_generate_v1() NOT NULL
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 174 (class 1259 OID 16404)
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- TOC entry 2070 (class 0 OID 0)
-- Dependencies: 174
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- TOC entry 1943 (class 2604 OID 16420)
-- Name: avatar id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.avatar ALTER COLUMN id SET DEFAULT nextval('public.avatar_id_seq'::regclass);


--
-- TOC entry 1941 (class 2604 OID 16409)
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- TOC entry 2061 (class 0 OID 16417)
-- Dependencies: 177
-- Data for Name: avatar; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.avatar (id, url, type) FROM stdin;
1	http://mydomain.com/user/user.jpg	image/jpeg
2	http://google.com	image/png
3	http://google.com	image/png
4	http://google.com	image/png
5	http://google.com	image/png
6	http://google.com	image/png
\.


--
-- TOC entry 2059 (class 0 OID 16406)
-- Dependencies: 175
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, name, email, password, cpf, avatar, datastart, uuiduser) FROM stdin;
1	Junior S3WF	email@email.com	senha123	999.999.999-99	1	2019-07-03	ba6cf43c-9d88-11e9-86bd-37b0fb8532e9
3	User Name	test@email.com	pwd123	888.888.888-88	6	07-03-2019	2c374a08-9d99-11e9-bbb7-3f8356707b46
\.


--
-- TOC entry 2071 (class 0 OID 0)
-- Dependencies: 176
-- Name: avatar_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.avatar_id_seq', 6, true);


--
-- TOC entry 2072 (class 0 OID 0)
-- Dependencies: 174
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 3, true);


--
-- TOC entry 1947 (class 2606 OID 16422)
-- Name: avatar avatar_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.avatar
    ADD CONSTRAINT avatar_pk PRIMARY KEY (id);


--
-- TOC entry 1945 (class 2606 OID 16414)
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);


--
-- TOC entry 1948 (class 2606 OID 16423)
-- Name: users users_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk0 FOREIGN KEY (avatar) REFERENCES public.avatar(id);


--
-- TOC entry 2067 (class 0 OID 0)
-- Dependencies: 7
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


-- Completed on 2019-07-03 10:50:53 -03

--
-- PostgreSQL database dump complete
--
