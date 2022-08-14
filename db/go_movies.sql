--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.genres (
    id integer NOT NULL,
    genre_name character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: genres_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.genres_id_seq OWNED BY public.genres.id;


--
-- Name: movies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying,
    description text,
    year integer,
    release_date date,
    runtime integer,
    rating integer,
    mpaa_rating character varying,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: movies_genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_genres (
    id integer NOT NULL,
    movie_id integer,
    genre_id integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);


--
-- Name: movies_genres_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.movies_genres_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: movies_genres_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.movies_genres_id_seq OWNED BY public.movies_genres.id;


--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: genres id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.genres ALTER COLUMN id SET DEFAULT nextval('public.genres_id_seq'::regclass);


--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: movies_genres id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_genres ALTER COLUMN id SET DEFAULT nextval('public.movies_genres_id_seq'::regclass);


--
-- Data for Name: genres; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.genres (id, genre_name, created_at, updated_at) FROM stdin;
1	Drama	2021-05-17 00:00:00	2021-05-17 00:00:00
2	Crime	2021-05-17 00:00:00	2021-05-17 00:00:00
3	Action	2021-05-17 00:00:00	2021-05-17 00:00:00
4	Comic Book	2021-05-17 00:00:00	2021-05-17 00:00:00
5	Sci-Fi	2021-05-17 00:00:00	2021-05-17 00:00:00
6	Mystery	2021-05-17 00:00:00	2021-05-17 00:00:00
7	Adventure	2021-05-17 00:00:00	2021-05-17 00:00:00
8	Comedy	2021-05-17 00:00:00	2021-05-17 00:00:00
9	Romance	2021-05-17 00:00:00	2021-05-17 00:00:00
\.


--
-- Data for Name: movies; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.movies (id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at) FROM stdin;
1	The Shawshank Redemption	Two imprisoned men bond over a number of years	1994	1994-10-14	142	5	R	2021-05-17 00:00:00	2021-05-17 00:00:00
2	The Godfather	The aging patriarch of an organized crime dynasty transfers control to his son	1972	1972-03-24	175	5	R	2021-05-17 00:00:00	2021-05-17 00:00:00
4	American Psycho	 A wealthy New York investment banking executive hides his alternate psychopathic ego	2000	2000-04-14	102	4	R	2021-05-17 00:00:00	2021-05-17 00:00:00
3	The Dark Knight	The menace known as the Joker wreaks havoc on Gotham City	2008	2008-07-18	152	5	PG13	2021-05-17 00:00:00	2021-05-17 00:00:00
\.


--
-- Data for Name: movies_genres; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.movies_genres (id, movie_id, genre_id, created_at, updated_at) FROM stdin;
\.


--
-- Name: genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.genres_id_seq', 9, true);


--
-- Name: movies_genres_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.movies_genres_id_seq', 1, false);


--
-- Name: movies_id_seq; Type: SEQUENCE SET; Schema: public; Owner: -
--

SELECT pg_catalog.setval('public.movies_id_seq', 4, true);


--
-- Name: genres genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.genres
    ADD CONSTRAINT genres_pkey PRIMARY KEY (id);


--
-- Name: movies_genres movies_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_pkey PRIMARY KEY (id);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: movies_genres fk_movie_genries_genre_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT fk_movie_genries_genre_id FOREIGN KEY (genre_id) REFERENCES public.genres(id);


--
-- Name: movies_genres fk_movie_genries_movie_id; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT fk_movie_genries_movie_id FOREIGN KEY (movie_id) REFERENCES public.movies(id);


--
-- PostgreSQL database dump complete
--

