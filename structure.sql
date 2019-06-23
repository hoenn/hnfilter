--
-- PostgreSQL database dump
--

-- Dumped from database version 10.8
-- Dumped by pg_dump version 11.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;


CREATE ROLE doadmin superuser;

--
-- Name: tsv_update_trigger(); Type: FUNCTION; Schema: public; Owner: doadmin
--

CREATE FUNCTION public.tsv_update_trigger() RETURNS trigger
    LANGUAGE plpgsql
    AS $$  
begin  
  new.tsv :=
     setweight(to_tsvector(new.body), 'A');
  return new;
end  
$$;


ALTER FUNCTION public.tsv_update_trigger() OWNER TO doadmin;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: doadmin
--

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
-- Name: comments update_tsv; Type: TRIGGER; Schema: public; Owner: doadmin
--

CREATE TRIGGER update_tsv BEFORE INSERT OR UPDATE ON public.comments FOR EACH ROW EXECUTE PROCEDURE public.tsv_update_trigger();


--
-- PostgreSQL database dump complete
--

