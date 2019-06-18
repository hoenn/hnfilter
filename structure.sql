--
-- PostgreSQL database dump
--

-- Dumped from database version 10.8
-- Dumped by pg_dump version 10.8 (Ubuntu 10.8-0ubuntu0.18.04.1)

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
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: doadmin
--

CREATE ROLE doadmin superuser;

CREATE TABLE public.comments (
    author character varying(255),
    id integer NOT NULL,
    kids integer[],
    parent integer,
    body text,
    "time" timestamp without time zone,
    tsv tsvector
);


ALTER TABLE public.comments OWNER TO doadmin;

--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: doadmin
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: idx_comments_tsv; Type: INDEX; Schema: public; Owner: doadmin
--

CREATE INDEX idx_comments_tsv ON public.comments USING gin (tsv);


--
-- Name: comments tsvectorupdate; Type: TRIGGER; Schema: public; Owner: doadmin
--

CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE ON public.comments FOR EACH ROW EXECUTE PROCEDURE tsvector_update_trigger('tsv', 'pg_catalog.english', 'body');


--
-- PostgreSQL database dump complete
--

