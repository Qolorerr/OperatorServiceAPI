--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

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
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: appeal_tag_links; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.appeal_tag_links (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    appeal_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.appeal_tag_links OWNER TO test_user;

--
-- Name: appeals; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.appeals (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    weight integer NOT NULL
);


ALTER TABLE public.appeals OWNER TO test_user;

--
-- Name: group_tag_links; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.group_tag_links (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    group_id uuid NOT NULL,
    tag_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.group_tag_links OWNER TO test_user;

--
-- Name: operator_group_links; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.operator_group_links (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    operator_id uuid NOT NULL,
    group_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.operator_group_links OWNER TO test_user;

--
-- Name: operator_groups; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.operator_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.operator_groups OWNER TO test_user;

--
-- Name: operators; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.operators (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.operators OWNER TO test_user;

--
-- Name: tags; Type: TABLE; Schema: public; Owner: test_user
--

CREATE TABLE public.tags (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    deleted boolean DEFAULT false,
    name text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tags OWNER TO test_user;

--
-- Data for Name: appeal_tag_links; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.appeal_tag_links (id, deleted, appeal_id, tag_id, created_at, updated_at) FROM stdin;
92d33d1c-2778-4232-84a7-5b4964d1c515	f	1db0c609-f0de-4ad9-a239-e98b5598c24e	4fec0a08-ab99-4d74-8f46-efafac1f51a5	2025-02-09 20:33:21.73421	2025-02-09 20:33:21.73421
046493ee-8364-4811-8b8e-e4e36f05c5b9	f	de85f2c7-d62c-43f7-88cf-f6632a615a5b	4fec0a08-ab99-4d74-8f46-efafac1f51a5	2025-02-09 20:58:21.711349	2025-02-09 20:58:21.711349
\.


--
-- Data for Name: appeals; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.appeals (id, deleted, user_id, created_at, updated_at, weight) FROM stdin;
8d88bd3a-7f34-42e8-b2ca-c28c66f9d234	f	23305c03-4153-4e3a-b7ba-363a203cb2e0	2025-02-09 20:24:06.855716	2025-02-09 20:24:06.855716	0
1db0c609-f0de-4ad9-a239-e98b5598c24e	f	23305c03-4153-4e3a-b7ba-363a203cb2e0	2025-02-09 20:33:21.733572	2025-02-09 20:33:21.733572	1
de85f2c7-d62c-43f7-88cf-f6632a615a5b	f	23305c03-4153-4e3a-b7ba-363a203cb2e0	2025-02-09 20:58:21.707929	2025-02-09 20:58:21.707929	1
\.


--
-- Data for Name: group_tag_links; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.group_tag_links (id, deleted, group_id, tag_id, created_at, updated_at) FROM stdin;
acae3873-eebf-46fa-a26e-4559331e1f82	f	ddac7502-44f4-460a-97f2-4bcc0c16bbca	4fec0a08-ab99-4d74-8f46-efafac1f51a5	2025-02-09 20:41:04.445342	2025-02-09 20:41:04.445342
b185a96e-3fff-4760-854c-83bbdb820468	f	ddac7502-44f4-460a-97f2-4bcc0c16bbca	8fc7501e-9bf2-4365-aa10-4e11132e181c	2025-02-09 20:41:04.446194	2025-02-09 20:41:04.446194
\.


--
-- Data for Name: operator_group_links; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.operator_group_links (id, deleted, operator_id, group_id, created_at, updated_at) FROM stdin;
60044ea8-8127-42f1-8797-b38de98574ee	f	6dd02caf-8b04-4b58-9e3b-0ac915f01bf6	ddac7502-44f4-460a-97f2-4bcc0c16bbca	2025-02-09 20:44:51.795174	2025-02-09 20:44:51.795174
\.


--
-- Data for Name: operator_groups; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.operator_groups (id, deleted, created_at, updated_at) FROM stdin;
ddac7502-44f4-460a-97f2-4bcc0c16bbca	f	2025-02-09 20:41:04.444585	2025-02-09 20:41:04.444585
\.


--
-- Data for Name: operators; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.operators (id, deleted, created_at, updated_at) FROM stdin;
6dd02caf-8b04-4b58-9e3b-0ac915f01bf6	f	2025-02-09 20:34:54.553694	2025-02-09 20:34:54.553694
\.


--
-- Data for Name: tags; Type: TABLE DATA; Schema: public; Owner: test_user
--

COPY public.tags (id, deleted, name, created_at, updated_at) FROM stdin;
4fec0a08-ab99-4d74-8f46-efafac1f51a5	f	first_tag	2025-02-09 20:15:47.736193	2025-02-09 20:15:47.736193
8fc7501e-9bf2-4365-aa10-4e11132e181c	f	second_tag	2025-02-09 20:16:04.706413	2025-02-09 20:16:04.706413
b1ed5d4a-a7a3-4840-a37c-4384977765f2	t	third_tag	2025-02-09 20:16:13.609063	2025-02-09 20:19:57.006263
\.


--
-- Name: appeal_tag_links appeal_tag_links_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.appeal_tag_links
    ADD CONSTRAINT appeal_tag_links_pkey PRIMARY KEY (id);


--
-- Name: appeals appeals_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.appeals
    ADD CONSTRAINT appeals_pkey PRIMARY KEY (id);


--
-- Name: group_tag_links group_tag_links_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.group_tag_links
    ADD CONSTRAINT group_tag_links_pkey PRIMARY KEY (id);


--
-- Name: operator_group_links operator_group_links_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.operator_group_links
    ADD CONSTRAINT operator_group_links_pkey PRIMARY KEY (id);


--
-- Name: operator_groups operator_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.operator_groups
    ADD CONSTRAINT operator_groups_pkey PRIMARY KEY (id);


--
-- Name: operators operators_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.operators
    ADD CONSTRAINT operators_pkey PRIMARY KEY (id);


--
-- Name: tags tags_pkey; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT tags_pkey PRIMARY KEY (id);


--
-- Name: tags uni_tags_name; Type: CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.tags
    ADD CONSTRAINT uni_tags_name UNIQUE (name);


--
-- Name: appeal_tag_links fk_appeals_appeal_tag_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.appeal_tag_links
    ADD CONSTRAINT fk_appeals_appeal_tag_links FOREIGN KEY (appeal_id) REFERENCES public.appeals(id);


--
-- Name: group_tag_links fk_operator_groups_group_tag_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.group_tag_links
    ADD CONSTRAINT fk_operator_groups_group_tag_links FOREIGN KEY (group_id) REFERENCES public.operator_groups(id);


--
-- Name: operator_group_links fk_operator_groups_operator_group_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.operator_group_links
    ADD CONSTRAINT fk_operator_groups_operator_group_links FOREIGN KEY (group_id) REFERENCES public.operator_groups(id);


--
-- Name: operator_group_links fk_operators_operator_group_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.operator_group_links
    ADD CONSTRAINT fk_operators_operator_group_links FOREIGN KEY (operator_id) REFERENCES public.operators(id);


--
-- Name: appeal_tag_links fk_tags_appeal_tag_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.appeal_tag_links
    ADD CONSTRAINT fk_tags_appeal_tag_links FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- Name: group_tag_links fk_tags_group_tag_links; Type: FK CONSTRAINT; Schema: public; Owner: test_user
--

ALTER TABLE ONLY public.group_tag_links
    ADD CONSTRAINT fk_tags_group_tag_links FOREIGN KEY (tag_id) REFERENCES public.tags(id);


--
-- PostgreSQL database dump complete
--

