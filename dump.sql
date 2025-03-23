--
-- Name: fulfillment_types; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.fulfillment_types (
    id bigint NOT NULL,
    modifier_id bigint,
    name character varying(55) NOT NULL,
    description character varying(155)
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
    name character varying(55) NOT NULL,
    modifier_id bigint,
    description character varying(155)
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
-- Name: order_price_modifiers; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.order_price_modifiers (
    id bigint NOT NULL,
    modifier_type_id bigint NOT NULL,
    description character varying(255) NOT NULL,
    percent double precision NOT NULL,
    order_id bigint,
    service_id bigint
);


ALTER TABLE public.order_price_modifiers OWNER TO pguser;

--
-- Name: order_price_modifiers_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.order_price_modifiers ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.order_price_modifiers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: orders; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    user_name character varying(55) NOT NULL,
    phone_number character varying(55) NOT NULL,
    total double precision NOT NULL,
    final double precision NOT NULL,
    creation_date timestamp without time zone NOT NULL
);


ALTER TABLE public.orders OWNER TO pguser;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.orders ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: orders_service_items; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.orders_service_items (
    id bigint NOT NULL,
    service_item_id bigint NOT NULL,
    quantity double precision NOT NULL,
    price double precision NOT NULL,
    order_service_id bigint NOT NULL
);


ALTER TABLE public.orders_service_items OWNER TO pguser;

--
-- Name: orders_service_items_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.orders_service_items ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_service_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: orders_services; Type: TABLE; Schema: public; Owner: pguser
--

CREATE TABLE public.orders_services (
    id bigint NOT NULL,
    order_id bigint NOT NULL,
    service_id bigint NOT NULL
);


ALTER TABLE public.orders_services OWNER TO pguser;

--
-- Name: orders_services_id_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

ALTER TABLE public.orders_services ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.orders_services_id_seq
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
    item_id bigint NOT NULL,
    price double precision NOT NULL,
    service_id bigint,
    sub_service_id bigint
);


ALTER TABLE public.service_items OWNER TO pguser;

--
-- Name: service_items_seq; Type: SEQUENCE; Schema: public; Owner: pguser
--

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
    name character varying(255) NOT NULL,
    unit_id bigint
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
    modifier_id bigint NOT NULL,
    unit_quantity double precision,
    description character varying(155)
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

COPY public.fulfillment_types (id, modifier_id, name, description) FROM stdin;
3	4	Экспресс	Наценка за экспресс выполнение
2	\N	Стандартная	\N
1	2	Неспешная	Скидка за неспешное выполнение
\.


--
-- Data for Name: item_types; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.item_types (id, name, modifier_id, description) FROM stdin;
2	Взрослая	\N	\N
1	Детская	3	Скидка за детские вещи
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
-- Data for Name: order_price_modifiers; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.order_price_modifiers (id, modifier_type_id, description, percent, order_id, service_id) FROM stdin;
7	1	Скидка за более 10кг вещей	20	\N	7
8	2	Наценка за экспресс выполнение	50	9	\N
9	1	Скидка за более 10кг вещей	20	\N	8
10	2	Наценка за экспресс выполнение	50	10	\N
11	1	Скидка за детские вещи	50	\N	11
12	1	Скидка за более 10кг вещей	20	\N	11
13	1	Скидка за детские вещи	50	\N	13
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.orders (id, user_name, phone_number, total, final, creation_date) FROM stdin;
9	Исмаил	+998998152821	141.8	170.1	2025-03-21 02:48:00
10	Исмаил	+998998152821	141.8	170.1	2025-03-21 21:49:12.402377
11	Test		40	40	2025-03-23 20:13:17.454124
12	TEst2		7	7	2025-03-23 20:14:44.335402
13	TEssadsa	41432534512321	99	39.6	2025-03-23 20:59:18.995144
14	TEste231	415434312423	93.5	73.5	2025-03-23 21:19:37.510935
\.


--
-- Data for Name: orders_service_items; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.orders_service_items (id, service_item_id, quantity, price, order_service_id) FROM stdin;
17	103	2	20	9
18	126	2	3.5	10
19	119	11	9	11
20	115	5	10	12
21	103	2	20	13
22	127	1	3.5	14
\.


--
-- Data for Name: orders_services; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.orders_services (id, order_id, service_id) FROM stdin;
7	9	3
8	10	3
9	11	1
10	12	5
11	13	3
12	14	3
13	14	1
14	14	5
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

COPY public.service_items (id, item_id, price, service_id, sub_service_id) FROM stdin;
103	1	20	1	\N
104	2	10	1	\N
105	3	15	1	\N
106	4	15	1	\N
107	5	15	1	\N
108	6	20	1	\N
109	1	5	2	\N
110	2	5	2	\N
111	3	5	2	\N
112	4	5	2	\N
113	5	5	2	\N
114	6	5	2	\N
115	7	10	3	\N
116	8	8	3	\N
117	9	12	3	\N
118	10	15	3	\N
119	11	9	3	\N
120	12	10	3	\N
121	13	5	4	\N
122	2	3	4	\N
123	14	2	4	\N
124	5	6	4	\N
125	3	8	4	\N
126	1	3.5	\N	1
127	2	3.5	\N	1
128	3	3.5	\N	1
129	4	3.5	\N	1
130	5	3.5	\N	1
131	6	3.5	\N	1
132	13	3.5	\N	1
133	14	3.5	\N	1
134	1	3.5	\N	2
135	2	3.5	\N	2
136	3	3.5	\N	2
137	4	3.5	\N	2
138	5	3.5	\N	2
139	6	3.5	\N	2
140	13	3.5	\N	2
141	14	3.5	\N	2
142	1	5	\N	3
143	2	5	\N	3
144	3	5	\N	3
145	4	5	\N	3
146	5	5	\N	3
147	6	5	\N	3
148	13	5	\N	3
149	14	5	\N	3
150	1	3	\N	4
151	2	3	\N	4
152	3	3	\N	4
153	4	3	\N	4
154	5	3	\N	4
155	6	3	\N	4
156	13	3	\N	4
157	14	3	\N	4
158	1	2	\N	5
159	2	2	\N	5
160	3	2	\N	5
161	4	2	\N	5
162	5	2	\N	5
163	6	2	\N	5
164	13	2	\N	5
165	14	2	\N	5
\.


--
-- Data for Name: services; Type: TABLE DATA; Schema: public; Owner: pguser
--

COPY public.services (id, name, unit_id) FROM stdin;
1	Химчистка	2
2	Ручная стирка	2
3	Общие услуги по стирке	1
4	Гладильные услуги	2
5	Ремонт одежды	2
6	Удаление пятен	2
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

COPY public.unit_modifiers (id, unit_id, modifier_id, unit_quantity, description) FROM stdin;
3	1	1	10	Скидка за более 10кг вещей
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
-- Name: order_price_modifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.order_price_modifiers_id_seq', 13, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.orders_id_seq', 14, true);


--
-- Name: orders_service_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.orders_service_items_id_seq', 22, true);


--
-- Name: orders_services_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.orders_services_id_seq', 14, true);


--
-- Name: price_modifiers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.price_modifiers_id_seq', 2, true);


--
-- Name: price_modifiers_id_seq1; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.price_modifiers_id_seq1', 4, true);


--
-- Name: service_items_seq; Type: SEQUENCE SET; Schema: public; Owner: pguser
--

SELECT pg_catalog.setval('public.service_items_seq', 165, true);


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

SELECT pg_catalog.setval('public.unit_modifiers_id_seq', 3, true);


--
-- Name: fulfillment_types fulfillment_types_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.fulfillment_types
    ADD CONSTRAINT fulfillment_types_pkey PRIMARY KEY (id);


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
-- Name: order_price_modifiers order_price_modifiers_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.order_price_modifiers
    ADD CONSTRAINT order_price_modifiers_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: orders_service_items orders_service_items_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_service_items
    ADD CONSTRAINT orders_service_items_pkey PRIMARY KEY (id);


--
-- Name: orders_services orders_services_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_services
    ADD CONSTRAINT orders_services_pkey PRIMARY KEY (id);


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
-- Name: service_items service_items_pkey; Type: CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_pkey PRIMARY KEY (id);


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
    ADD CONSTRAINT fulfillment_types_price_modifier_id_fkey FOREIGN KEY (modifier_id) REFERENCES public.price_modifiers(id) NOT VALID;


--
-- Name: item_types item_types_price_modifier_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.item_types
    ADD CONSTRAINT item_types_price_modifier_id_fkey FOREIGN KEY (modifier_id) REFERENCES public.price_modifiers(id);


--
-- Name: order_price_modifiers order_price_modifiers_modifier_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.order_price_modifiers
    ADD CONSTRAINT order_price_modifiers_modifier_type_id_fkey FOREIGN KEY (modifier_type_id) REFERENCES public.price_modifiers_types(id);


--
-- Name: order_price_modifiers order_price_modifiers_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.order_price_modifiers
    ADD CONSTRAINT order_price_modifiers_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) NOT VALID;


--
-- Name: order_price_modifiers order_price_modifiers_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.order_price_modifiers
    ADD CONSTRAINT order_price_modifiers_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.orders_services(id) NOT VALID;


--
-- Name: orders_service_items orders_service_items_order_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_service_items
    ADD CONSTRAINT orders_service_items_order_service_id_fkey FOREIGN KEY (order_service_id) REFERENCES public.orders_services(id) NOT VALID;


--
-- Name: orders_service_items orders_service_items_service_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_service_items
    ADD CONSTRAINT orders_service_items_service_item_id_fkey FOREIGN KEY (service_item_id) REFERENCES public.service_items(id);


--
-- Name: orders_services orders_services_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_services
    ADD CONSTRAINT orders_services_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id);


--
-- Name: orders_services orders_services_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_services
    ADD CONSTRAINT orders_services_service_id_fkey FOREIGN KEY (service_id) REFERENCES public.services(id) NOT VALID;


--
-- Name: orders_services orders_services_service_id_fkey1; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.orders_services
    ADD CONSTRAINT orders_services_service_id_fkey1 FOREIGN KEY (service_id) REFERENCES public.sub_services(id) NOT VALID;


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
-- Name: service_items service_items_sub_service_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.service_items
    ADD CONSTRAINT service_items_sub_service_id_fkey FOREIGN KEY (sub_service_id) REFERENCES public.sub_services(id) NOT VALID;


--
-- Name: services services_unit_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: pguser
--

ALTER TABLE ONLY public.services
    ADD CONSTRAINT services_unit_id_fkey FOREIGN KEY (unit_id) REFERENCES public.unit(id) NOT VALID;


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

