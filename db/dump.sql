--
-- PostgreSQL database dump
--

-- Dumped from database version 12.5
-- Dumped by pg_dump version 12.5

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


--
-- Name: add_item_to_user(character varying, character varying, integer); Type: PROCEDURE; Schema: public; Owner: usanin
--

CREATE PROCEDURE public.add_item_to_user(target_item_id character varying, target_user_id character varying, target_count integer)
    LANGUAGE plpgsql
    AS $$
declare
    target_game_id varchar(36);
    target_inventory_id varchar(36);
begin
    target_game_id = (select game_id from item where id=target_item_id);

    insert into inventory(user_id, game_id) values(target_user_id, target_game_id) on conflict do nothing;
    target_inventory_id = (select i.id from inventory i where i.user_id=target_user_id and i.game_id=target_game_id);

    update inventory_item set count=count+target_count where inventory_id = target_inventory_id and item_id =target_item_id;
    if not found then
        insert into inventory_item(inventory_id, item_id, count) values (target_inventory_id, target_item_id, target_count);
    end if;

    commit;
end;
$$;


ALTER PROCEDURE public.add_item_to_user(target_item_id character varying, target_user_id character varying, target_count integer) OWNER TO usanin;

--
-- Name: close_lot(character varying); Type: PROCEDURE; Schema: public; Owner: usanin
--

CREATE PROCEDURE public.close_lot(target_lot_id character varying)
    LANGUAGE plpgsql
    AS $$
declare
    won_bet_id varchar(36);
    won_bet_user_id varchar(36);
    lot_owner_user_id varchar(36);
    won_bet_price real;
begin
    won_bet_id = (select id from bet b where b.lot_id=target_lot_id ORDER BY total_price DESC LIMIT 1);
    won_bet_price = (select total_price from bet where id=won_bet_id);
    won_bet_user_id = (select user_id from bet where id=won_bet_id);
    lot_owner_user_id = (select user_id from lot where id=target_lot_id);

    update "user" set bill = bill - won_bet_price where id=won_bet_user_id;
    update "user" set bill = bill + won_bet_price where id=lot_owner_user_id;

    commit;
end;$$;


ALTER PROCEDURE public.close_lot(target_lot_id character varying) OWNER TO usanin;

--
-- Name: deactivate_old_emails(); Type: FUNCTION; Schema: public; Owner: usanin
--

CREATE FUNCTION public.deactivate_old_emails() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.is_active > 0 THEN
        UPDATE email SET is_active=0 WHERE user_id=NEW.user_id AND id <> NEW.id;
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.deactivate_old_emails() OWNER TO usanin;

--
-- Name: update_max_bet_price(); Type: FUNCTION; Schema: public; Owner: usanin
--

CREATE FUNCTION public.update_max_bet_price() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF NEW.total_price > (SELECT max_bet_price FROM lot WHERE id=NEW.lot_id) THEN
        UPDATE lot SET max_bet_price=NEW.total_price WHERE id=NEW.lot_id;
    END IF;

    RETURN NEW;
END;
$$;


ALTER FUNCTION public.update_max_bet_price() OWNER TO usanin;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: bet; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.bet (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    lot_id character varying(36) NOT NULL,
    user_id character varying(36) NOT NULL,
    total_price real DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_DATE NOT NULL
);


ALTER TABLE public.bet OWNER TO usanin;

--
-- Name: email; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.email (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    email character varying(50) NOT NULL,
    user_id character varying(36) NOT NULL,
    is_active integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.email OWNER TO usanin;

--
-- Name: game; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.game (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.game OWNER TO usanin;

--
-- Name: inventory; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.inventory (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id character varying(36) NOT NULL,
    game_id character varying(36) NOT NULL
);


ALTER TABLE public.inventory OWNER TO usanin;

--
-- Name: inventory_item; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.inventory_item (
    inventory_id character varying(36) NOT NULL,
    item_id character varying(36) NOT NULL,
    count integer NOT NULL
);


ALTER TABLE public.inventory_item OWNER TO usanin;

--
-- Name: item; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.item (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying(255) NOT NULL,
    image_url character varying(255),
    game_id character varying(255),
    description character varying(255),
    data jsonb
);


ALTER TABLE public.item OWNER TO usanin;

--
-- Name: lot; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.lot (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id character varying(36) NOT NULL,
    begin_price real DEFAULT 0 NOT NULL,
    buy_price real,
    created_at timestamp without time zone DEFAULT CURRENT_DATE NOT NULL,
    end_at timestamp without time zone,
    is_closed integer DEFAULT 0 NOT NULL,
    max_bet_price real DEFAULT 0 NOT NULL
);


ALTER TABLE public.lot OWNER TO usanin;

--
-- Name: lot_item; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.lot_item (
    lot_id character varying(36) NOT NULL,
    item_id character varying(36) NOT NULL,
    count integer
);


ALTER TABLE public.lot_item OWNER TO usanin;

--
-- Name: phone; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public.phone (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    phone character varying(50) NOT NULL,
    user_id character varying(36) NOT NULL,
    is_active integer DEFAULT 0 NOT NULL
);


ALTER TABLE public.phone OWNER TO usanin;

--
-- Name: user; Type: TABLE; Schema: public; Owner: usanin
--

CREATE TABLE public."user" (
    id character varying(36) DEFAULT public.uuid_generate_v4() NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_DATE,
    last_login_at timestamp without time zone DEFAULT CURRENT_DATE,
    bill double precision DEFAULT 0 NOT NULL
);


ALTER TABLE public."user" OWNER TO usanin;

--
-- Data for Name: bet; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.bet (id, lot_id, user_id, total_price, created_at) FROM stdin;
60f9671d-3ac8-4bc7-91c3-0e95186a77bf	e4abbf77-0359-49ae-a945-03d0f2e573fb	30eff35d-46c2-4483-874f-3607e269447e	500	2020-12-12 00:00:00
a4779dbe-1e7f-47df-85d8-83babdaf5d0d	e3aba2d2-47c9-440e-8e5b-cf753408a06b	e219240d-7b25-4f4a-8b98-c4f90f9b6790	3000	2020-12-12 00:00:00
f8e1c69a-67c6-4462-ae48-5988b99055b0	e3aba2d2-47c9-440e-8e5b-cf753408a06b	ce0d8279-08cd-438a-9949-a740aca9592f	5000	2020-12-12 00:00:00
fca2e632-41f9-483f-98b2-f8f1fa7736db	aa9efb6a-23ae-48c6-b17e-f97106d92f98	30eff35d-46c2-4483-874f-3607e269447e	1000	2020-12-12 00:00:00
2ed1915b-fab1-4ccf-b57c-cd993c9828e9	e3aba2d2-47c9-440e-8e5b-cf753408a06b	e219240d-7b25-4f4a-8b98-c4f90f9b6790	6000	2020-12-12 00:00:00
\.


--
-- Data for Name: email; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.email (id, email, user_id, is_active) FROM stdin;
ba304f5d-d71d-4a09-b40c-b228c6ffbb00	test@mail.ru	e219240d-7b25-4f4a-8b98-c4f90f9b6790	1
5642fd28-8951-478d-ae21-36ab4fbc6f4d	test2@mail.com	ce0d8279-08cd-438a-9949-a740aca9592f	1
74d9c9c6-8846-4f2e-a5ea-f33b10247e59	hidan9812@gmail.com	30eff35d-46c2-4483-874f-3607e269447e	1
009853e6-6bcb-4165-811a-db818142872a	speaker@mail.com	30eff35d-46c2-4483-874f-3607e269447e	0
\.


--
-- Data for Name: game; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.game (id, name) FROM stdin;
923b1c32-e00d-415e-9f9e-f83a2b07d52e	Blade and soul
e400efb5-68c5-4dc6-bb4a-e163bac833dc	Black Desert
fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	Lineage 2
b142a159-5cec-4d3b-a3d8-3b23de2564d8	World of Warcraft
\.


--
-- Data for Name: inventory; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.inventory (id, user_id, game_id) FROM stdin;
caa00108-bca4-4acf-ae1c-c4c6bf239276	30eff35d-46c2-4483-874f-3607e269447e	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932
4954ea64-dd45-42cf-9226-55782bab4b64	30eff35d-46c2-4483-874f-3607e269447e	b142a159-5cec-4d3b-a3d8-3b23de2564d8
8cd38cea-c6b7-4491-9175-e14c8ec885bf	e219240d-7b25-4f4a-8b98-c4f90f9b6790	923b1c32-e00d-415e-9f9e-f83a2b07d52e
9c89cba7-46cc-498c-aead-3cffe21a02de	e219240d-7b25-4f4a-8b98-c4f90f9b6790	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932
f9bf12db-f9e2-47ee-9b91-503a0404994b	e219240d-7b25-4f4a-8b98-c4f90f9b6790	b142a159-5cec-4d3b-a3d8-3b23de2564d8
153de3ee-76b8-4a87-b001-0df4fbbdfeb8	ce0d8279-08cd-438a-9949-a740aca9592f	923b1c32-e00d-415e-9f9e-f83a2b07d52e
932057cc-f9d1-47a8-9e33-e2190cfb6c5e	ce0d8279-08cd-438a-9949-a740aca9592f	b142a159-5cec-4d3b-a3d8-3b23de2564d8
\.


--
-- Data for Name: inventory_item; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.inventory_item (inventory_id, item_id, count) FROM stdin;
caa00108-bca4-4acf-ae1c-c4c6bf239276	046b5be3-5aef-4407-8bf5-019409011f71	2
caa00108-bca4-4acf-ae1c-c4c6bf239276	8d640670-d4e8-4b27-b6c1-7e1cfc8748f5	1
f9bf12db-f9e2-47ee-9b91-503a0404994b	5de914cf-952a-4f63-a83e-dc8359ba48e3	2
9c89cba7-46cc-498c-aead-3cffe21a02de	046b5be3-5aef-4407-8bf5-019409011f71	1
9c89cba7-46cc-498c-aead-3cffe21a02de	8d640670-d4e8-4b27-b6c1-7e1cfc8748f5	3
932057cc-f9d1-47a8-9e33-e2190cfb6c5e	c20827d4-068d-40fd-8a02-560b8041c457	1
caa00108-bca4-4acf-ae1c-c4c6bf239276	ccdf8f74-8d3d-42d2-bd13-c98c9ef34834	2
caa00108-bca4-4acf-ae1c-c4c6bf239276	c466a9bc-bc76-49df-a72f-768d75202f11	6
\.


--
-- Data for Name: item; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.item (id, name, image_url, game_id, description, data) FROM stdin;
ccdf8f74-8d3d-42d2-bd13-c98c9ef34834	Imperial Staff	https://l2db.ru/themes/l2db/images/items/weapon_imperial_staff_i01_0.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "blant", "attack": 274}
046b5be3-5aef-4407-8bf5-019409011f71	Meteor Shower	https://l2db.ru/themes/l2db/images/items/weapon_meteor_shower_i01.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "blant", "attack": 213}
c466a9bc-bc76-49df-a72f-768d75202f11	Arcana Mace	https://l2db.ru/themes/l2db/images/items/weapon_arcana_mace_i01_0.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "blant", "attack": 225}
8d640670-d4e8-4b27-b6c1-7e1cfc8748f5	Draconic Bow	https://l2db.ru/themes/l2db/images/items/weapon_draconic_bow_i00_0.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "bow", "attack": 581}
df744b54-a4ea-4113-964a-f56acea51d63	Heaven's Divider	https://l2db.ru/themes/l2db/images/items/weapon_heavens_divider_i00_0.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "sword", "attack": 342}
c4939841-2f6e-4866-8b85-e472fa92c5fe	Dragon Slayer	https://l2db.ru/themes/l2db/images/items/weapon_dragon_slayer_i01.png	fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932	\N	{"type": "sword", "attack": 282}
5de914cf-952a-4f63-a83e-dc8359ba48e3	Keepcrawler's Gutripper	https://wow.zamimg.com/modelviewer/live/webthumbs/item/168/183720.webp	b142a159-5cec-4d3b-a3d8-3b23de2564d8	\N	{"type": "dagger"}
c20827d4-068d-40fd-8a02-560b8041c457	Sunblade	https://wow.zamimg.com/uploads/screenshots/small/430489.jpg	b142a159-5cec-4d3b-a3d8-3b23de2564d8	\N	{"type": "dagger"}
\.


--
-- Data for Name: lot; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.lot (id, user_id, begin_price, buy_price, created_at, end_at, is_closed, max_bet_price) FROM stdin;
e4abbf77-0359-49ae-a945-03d0f2e573fb	ce0d8279-08cd-438a-9949-a740aca9592f	200	5000	2020-12-10 00:00:00	2020-12-11 00:00:00	1	500
aa9efb6a-23ae-48c6-b17e-f97106d92f98	ce0d8279-08cd-438a-9949-a740aca9592f	500	1000	2020-12-12 00:00:00	2020-12-19 00:00:00	0	1000
e3aba2d2-47c9-440e-8e5b-cf753408a06b	30eff35d-46c2-4483-874f-3607e269447e	100	15000	2020-12-12 00:00:00	2020-12-26 00:00:00	0	6000
\.


--
-- Data for Name: lot_item; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.lot_item (lot_id, item_id, count) FROM stdin;
e3aba2d2-47c9-440e-8e5b-cf753408a06b	ccdf8f74-8d3d-42d2-bd13-c98c9ef34834	1
e4abbf77-0359-49ae-a945-03d0f2e573fb	8d640670-d4e8-4b27-b6c1-7e1cfc8748f5	1
aa9efb6a-23ae-48c6-b17e-f97106d92f98	8d640670-d4e8-4b27-b6c1-7e1cfc8748f5	1
e3aba2d2-47c9-440e-8e5b-cf753408a06b	046b5be3-5aef-4407-8bf5-019409011f71	1
\.


--
-- Data for Name: phone; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public.phone (id, phone, user_id, is_active) FROM stdin;
80347da1-a04c-451d-9f17-a00ea3a1f70e	+79125947072	30eff35d-46c2-4483-874f-3607e269447e	1
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: usanin
--

COPY public."user" (id, username, password, created_at, last_login_at, bill) FROM stdin;
e219240d-7b25-4f4a-8b98-c4f90f9b6790	twentysixth	111111	2020-12-12 00:00:00	2020-12-12 00:00:00	8000
30eff35d-46c2-4483-874f-3607e269447e	speaker	123456	2020-12-12 00:00:00	2020-12-12 00:00:00	9500
ce0d8279-08cd-438a-9949-a740aca9592f	avgust	123456	2020-12-12 00:00:00	2020-12-12 00:00:00	15500
\.


--
-- Name: bet bet_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.bet
    ADD CONSTRAINT bet_pk PRIMARY KEY (id);


--
-- Name: email email_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.email
    ADD CONSTRAINT email_pk PRIMARY KEY (id);


--
-- Name: game game_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.game
    ADD CONSTRAINT game_pk PRIMARY KEY (id);


--
-- Name: inventory_item inventory_item_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory_item
    ADD CONSTRAINT inventory_item_pk PRIMARY KEY (inventory_id, item_id);


--
-- Name: inventory inventory_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_pk PRIMARY KEY (id);


--
-- Name: item item_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.item
    ADD CONSTRAINT item_pk PRIMARY KEY (id);


--
-- Name: lot_item lot_item_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.lot_item
    ADD CONSTRAINT lot_item_pk PRIMARY KEY (lot_id, item_id);


--
-- Name: lot lot_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.lot
    ADD CONSTRAINT lot_pk PRIMARY KEY (id);


--
-- Name: phone phone_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.phone
    ADD CONSTRAINT phone_pk PRIMARY KEY (id);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (id);


--
-- Name: bet_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX bet_id_uindex ON public.bet USING btree (id);


--
-- Name: email_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX email_id_uindex ON public.email USING btree (id);


--
-- Name: game_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX game_id_uindex ON public.game USING btree (id);


--
-- Name: game_name_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX game_name_uindex ON public.game USING btree (name);


--
-- Name: inventory_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX inventory_id_uindex ON public.inventory USING btree (id);


--
-- Name: inventory_user_id_game_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX inventory_user_id_game_id_uindex ON public.inventory USING btree (user_id, game_id);


--
-- Name: item_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX item_id_uindex ON public.item USING btree (id);


--
-- Name: item_name_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX item_name_uindex ON public.item USING btree (name);


--
-- Name: lot_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX lot_id_uindex ON public.lot USING btree (id);


--
-- Name: phone_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX phone_id_uindex ON public.phone USING btree (id);


--
-- Name: user_id_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX user_id_uindex ON public."user" USING btree (id);


--
-- Name: user_name_uindex; Type: INDEX; Schema: public; Owner: usanin
--

CREATE UNIQUE INDEX user_name_uindex ON public."user" USING btree (username);


--
-- Name: email deactivate_old_emails_on_insert; Type: TRIGGER; Schema: public; Owner: usanin
--

CREATE TRIGGER deactivate_old_emails_on_insert AFTER INSERT ON public.email FOR EACH ROW EXECUTE FUNCTION public.deactivate_old_emails();


--
-- Name: email deactivate_old_emails_on_update; Type: TRIGGER; Schema: public; Owner: usanin
--

CREATE TRIGGER deactivate_old_emails_on_update AFTER UPDATE ON public.email FOR EACH ROW EXECUTE FUNCTION public.deactivate_old_emails();


--
-- Name: bet update_max_bet_price; Type: TRIGGER; Schema: public; Owner: usanin
--

CREATE TRIGGER update_max_bet_price AFTER INSERT ON public.bet FOR EACH ROW EXECUTE FUNCTION public.update_max_bet_price();


--
-- Name: bet bet_lot_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.bet
    ADD CONSTRAINT bet_lot_id_fk FOREIGN KEY (lot_id) REFERENCES public.lot(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: bet bet_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.bet
    ADD CONSTRAINT bet_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: email email_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.email
    ADD CONSTRAINT email_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: inventory inventory_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_game_id_fk FOREIGN KEY (game_id) REFERENCES public.game(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: inventory_item inventory_item_inventory_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory_item
    ADD CONSTRAINT inventory_item_inventory_id_fk FOREIGN KEY (inventory_id) REFERENCES public.inventory(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: inventory_item inventory_item_item_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory_item
    ADD CONSTRAINT inventory_item_item_id_fk FOREIGN KEY (item_id) REFERENCES public.item(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: inventory inventory_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.inventory
    ADD CONSTRAINT inventory_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: item item_game_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.item
    ADD CONSTRAINT item_game_id_fk FOREIGN KEY (game_id) REFERENCES public.game(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: lot_item lot_item_item_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.lot_item
    ADD CONSTRAINT lot_item_item_id_fk FOREIGN KEY (item_id) REFERENCES public.item(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: lot_item lot_item_lot_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.lot_item
    ADD CONSTRAINT lot_item_lot_id_fk FOREIGN KEY (lot_id) REFERENCES public.lot(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: lot lot_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.lot
    ADD CONSTRAINT lot_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: phone phone_user_id_fk; Type: FK CONSTRAINT; Schema: public; Owner: usanin
--

ALTER TABLE ONLY public.phone
    ADD CONSTRAINT phone_user_id_fk FOREIGN KEY (user_id) REFERENCES public."user"(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

