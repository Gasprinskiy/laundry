CREATE TABLE public.fulfillment_types (
    id bigint NOT NULL,
    price_modifier_id bigint,
    name character varying(55) NOT NULL
);


ALTER TABLE public.fulfillment_types OWNER TO pguser;

--
-- Name: fulfillment_types_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.fulfillment_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.fulfillment_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: item_types; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.item_types (
    id bigint NOT NULL,
    type character varying(55) NOT NULL,
    price_modifier_id bigint
);


ALTER TABLE public.item_types OWNER TO pguser;

--
-- Name: item_types_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.item_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.item_types_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: items; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.items (
    id bigint NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.items OWNER TO pguser;

--
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: price_modifiers; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.price_modifiers (
    id bigint NOT NULL,
    percent double precision NOT NULL,
    modifier_type_id bigint NOT NULL
);


ALTER TABLE public.price_modifiers OWNER TO pguser;

--
-- Name: price_modifiers_types; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.price_modifiers_types (
    id bigint NOT NULL,
    modifier character varying(50) NOT NULL
);


ALTER TABLE public.price_modifiers_types OWNER TO pguser;

--
-- Name: price_modifiers_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.price_modifiers_types ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.price_modifiers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: price_modifiers_id_seq1; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.price_modifiers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.price_modifiers_id_seq1
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: service_items; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.service_items (
    id bigint NOT NULL,
    service_id bigint NOT NULL,
    item_id bigint NOT NULL,
    service_unit_id bigint NOT NULL,
    price double precision NOT NULL,
    sub_service_id bigint
);


ALTER TABLE public.service_items OWNER TO pguser;


ALTER TABLE public.service_items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.service_items_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);

--
-- Name: unit; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.unit (
    id bigint NOT NULL,
    unit character varying(50) NOT NULL
);


ALTER TABLE public.unit OWNER TO pguser;

--
-- Name: service_unit_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.unit ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.service_unit_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: services; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.services (
    id bigint NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.services OWNER TO pguser;

--
-- Name: services_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.services ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: sub_services; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.sub_services (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    service_id bigint NOT NULL
);


ALTER TABLE public.sub_services OWNER TO pguser;

--
-- Name: sub_services_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.sub_services ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sub_services_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: unit_modifiers; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.unit_modifiers (
    id bigint NOT NULL,
    unit_id bigint NOT NULL,
    modifier_id bigint NOT NULL
);


ALTER TABLE public.unit_modifiers OWNER TO pguser;

--
-- Name: unit_modifiers_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.unit_modifiers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.unit_modifiers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: fulfillment_types; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.fulfillment_types (id, price_modifier_id, name) FROM stdin;
1	2	Неспешная
2	\N	Стандартная
3	4	Экспресс
\.


--
-- Data for Name: item_types; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.item_types (id, type, price_modifier_id) FROM stdin;
1	Детская	3
2	Взрослая	\N
\.


--
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.items (id, name) FROM stdin;
1	Пальто
2	Брюки
3	Костюм
4	Сюртук
6	Свадебное Платье
7	Белые
8	Цветные
9	Шерсть
10	Шёлк
11	Мягкие Игрушки
12	Постельное Бельё
5	Платье
13	Рубашка
14	Юбка
\.


--
-- Data for Name: price_modifiers; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.price_modifiers (id, percent, modifier_type_id) FROM stdin;
1	20	1
2	30	1
3	50	1
4	50	2
\.


--
-- Data for Name: price_modifiers_types; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.price_modifiers_types (id, modifier) FROM stdin;
1	discount
2	markup
\.


--
-- Data for Name: service_items; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.service_items (service_id, item_id, service_unit_id, price, sub_service_id) FROM stdin;
1	1	2	20	\N
1	2	2	10	\N
1	3	2	15	\N
1	4	2	15	\N
1	5	2	15	\N
1	6	2	20	\N
2	1	2	5	\N
2	2	2	5	\N
2	3	2	5	\N
2	4	2	5	\N
2	5	2	5	\N
2	6	2	5	\N
3	7	1	10	\N
3	8	1	8	\N
3	9	1	12	\N
3	10	1	15	\N
3	11	1	9	\N
3	12	1	10	\N
4	13	2	5	\N
4	2	2	3	\N
4	14	2	2	\N
4	5	2	6	\N
4	3	2	8	\N
5	1	2	3.5	1
5	2	2	3.5	1
5	3	2	3.5	1
5	4	2	3.5	1
5	5	2	3.5	1
5	6	2	3.5	1
5	13	2	3.5	1
5	14	2	3.5	1
5	1	2	3.5	2
5	2	2	3.5	2
5	3	2	3.5	2
5	4	2	3.5	2
5	5	2	3.5	2
5	6	2	3.5	2
5	13	2	3.5	2
5	14	2	3.5	2
6	1	2	5	3
6	2	2	5	3
6	3	2	5	3
6	4	2	5	3
6	5	2	5	3
6	6	2	5	3
6	13	2	5	3
6	14	2	5	3
6	1	2	3	4
6	2	2	3	4
6	3	2	3	4
6	4	2	3	4
6	5	2	3	4
6	6	2	3	4
6	13	2	3	4
6	14	2	3	4
6	1	2	2	5
6	2	2	2	5
6	3	2	2	5
6	4	2	2	5
6	5	2	2	5
6	6	2	2	5
6	13	2	2	5
6	14	2	2	5
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.services (id, name) FROM stdin;
1	Химчистка
2	Ручная стирка
3	Общие услуги по стирке
4	Гладильные услуги
5	Ремонт одежды
6	Удаление пятен
\.


--
-- Data for Name: sub_services; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.sub_services (id, name, service_id) FROM stdin;
1	Исправление шва	5
2	Исправление штопки	5
3	Пятна от масла	6
4	Пятна от крови	6
5	Общая грязь	6
\.


--
-- Data for Name: unit; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.unit (id, unit) FROM stdin;
1	kg
2	pcs
\.


--
-- Data for Name: unit_modifiers; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.unit_modifiers (id, unit_id, modifier_id) FROM stdin;
1	1	1
\.


--
-- Name: fulfillment_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.fulfillment_types_id_seq', 3, true);


--
-- Name: item_types_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.item_types_id_seq', 2, true);


--
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.items_id_seq', 19, true);


--
-- Name: price_modifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.price_modifiers_id_seq', 2, true);


--
-- Name: price_modifiers_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.price_modifiers_id_seq1', 4, true);

SELECT pg_catalog.setval('public.service_items_seq', 54, true);


--
-- Name: service_unit_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.service_unit_id_seq', 2, true);


--
-- Name: services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.services_id_seq', 6, true);


--
-- Name: sub_services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.sub_services_id_seq', 7, true);


--
-- Name: unit_modifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.unit_modifiers_id_seq', 1, true);


--
-- Name: item_types item_types_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.item_types
    ADD CONSTRAINT item_types_pkey PRIMARY KEY (id);


--
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- Name: price_modifiers_types price_modifiers_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.price_modifiers_types
    ADD CONSTRAINT price_modifiers_pkey PRIMARY KEY (id);


--
-- Name: price_modifiers price_modifiers_pkey1; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.price_modifiers
    ADD CONSTRAINT price_modifiers_pkey1 PRIMARY KEY (id);


--
-- Name: unit service_unit_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.unit
    ADD CONSTRAINT service_unit_pkey PRIMARY KEY (id);


--
-- Name: services services_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_pkey PRIMARY KEY (id);


--
-- Name: sub_services sub_services_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.sub_services
    ADD CONSTRAINT sub_services_pkey PRIMARY KEY (id);


--
-- Name: unit_modifiers unit_modifiers_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.unit_modifiers
    ADD CONSTRAINT unit_modifiers_pkey PRIMARY KEY (id);


--
-- Name: fulfillment_types fulfillment_types_price_modifier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.fulfillment_types
    ADD CONSTRAINT fulfillment_types_price_modifier_id_fkey FOREIGN KEY (price_modifier_id) REFERENCES public.price_modifiers(id) NOT VALID;


--
-- Name: item_types item_types_price_modifier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.item_types
    ADD CONSTRAINT item_types_price_modifier_id_fkey FOREIGN KEY (price_modifier_id) REFERENCES public.price_modifiers(id);


--
-- Name: price_modifiers price_modifiers_modifier_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.price_modifiers
    ADD CONSTRAINT price_modifiers_modifier_type_id_fkey FOREIGN KEY (modifier_type_id) REFERENCES public.price_modifiers_types(id);


--
-- Name: service_items service_items_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_item_id_fkey FOREIGN KEY (item_id) REFERENCES public.items(id);


--
-- Name: service_items service_items_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.services(id) NOT VALID;


--
-- Name: service_items service_items_service_unit_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_service_unit_id_fkey FOREIGN KEY (service_unit_id) REFERENCES public.unit(id) NOT VALID;


--
-- Name: service_items service_items_sub_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_sub_service_id_fkey FOREIGN KEY (sub_service_id) REFERENCES public.sub_services(id) NOT VALID;


--
-- Name: sub_services sub_services_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.sub_services
    ADD CONSTRAINT sub_services_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.services(id);


--
-- Name: unit_modifiers unit_modifiers_modifier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.unit_modifiers
    ADD CONSTRAINT unit_modifiers_modifier_id_fkey FOREIGN KEY (modifier_id) REFERENCES public.price_modifiers(id);


--
-- Name: unit_modifiers unit_modifiers_unit_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.unit_modifiers
    ADD CONSTRAINT unit_modifiers_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES public.unit(id);


--
-- PostgreSQL database dump complete
--

--
-- Database "postgres" dump
--

--
-- PostgreSQL database dump
--

-- Dumped from database version 15.12
-- Dumped by pg_dump version 15.12

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

DROP DATABASE postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: pguser
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE postgres OWNER TO pguser;

\connect postgres

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
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: pguser
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- PostgreSQL database dump complete
--

--
-- PostgreSQL database cluster dump complete
--