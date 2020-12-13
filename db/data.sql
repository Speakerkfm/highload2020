INSERT INTO public.game (id, name) VALUES ('923b1c32-e00d-415e-9f9e-f83a2b07d52e', 'Blade and soul');
INSERT INTO public.game (id, name) VALUES ('e400efb5-68c5-4dc6-bb4a-e163bac833dc', 'Black Desert');
INSERT INTO public.game (id, name) VALUES ('fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', 'Lineage 2');
INSERT INTO public.game (id, name) VALUES ('b142a159-5cec-4d3b-a3d8-3b23de2564d8', 'World of Warcraft');

INSERT INTO public."user" (id, username, password, created_at, last_login_at, bill) VALUES ('e219240d-7b25-4f4a-8b98-c4f90f9b6790', 'twentysixth', '111111', '2020-12-12 00:00:00.000000', '2020-12-12 00:00:00.000000', 8000);
INSERT INTO public."user" (id, username, password, created_at, last_login_at, bill) VALUES ('30eff35d-46c2-4483-874f-3607e269447e', 'speaker', '123456', '2020-12-12 00:00:00.000000', '2020-12-12 00:00:00.000000', 9500);
INSERT INTO public."user" (id, username, password, created_at, last_login_at, bill) VALUES ('ce0d8279-08cd-438a-9949-a740aca9592f', 'avgust', '123456', '2020-12-12 00:00:00.000000', '2020-12-12 00:00:00.000000', 15500);

INSERT INTO public.phone (id, phone, user_id, is_active) VALUES ('80347da1-a04c-451d-9f17-a00ea3a1f70e', '+79125947072', '30eff35d-46c2-4483-874f-3607e269447e', 1);

INSERT INTO public.email (id, email, user_id, is_active) VALUES ('ba304f5d-d71d-4a09-b40c-b228c6ffbb00', 'test@mail.ru', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', 1);
INSERT INTO public.email (id, email, user_id, is_active) VALUES ('5642fd28-8951-478d-ae21-36ab4fbc6f4d', 'test2@mail.com', 'ce0d8279-08cd-438a-9949-a740aca9592f', 1);
INSERT INTO public.email (id, email, user_id, is_active) VALUES ('74d9c9c6-8846-4f2e-a5ea-f33b10247e59', 'hidan9812@gmail.com', '30eff35d-46c2-4483-874f-3607e269447e', 1);
INSERT INTO public.email (id, email, user_id, is_active) VALUES ('009853e6-6bcb-4165-811a-db818142872a', 'speaker@mail.com', '30eff35d-46c2-4483-874f-3607e269447e', 0);

INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('ccdf8f74-8d3d-42d2-bd13-c98c9ef34834', 'Imperial Staff', 'https://l2db.ru/themes/l2db/images/items/weapon_imperial_staff_i01_0.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "blant", "attack": 274}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('046b5be3-5aef-4407-8bf5-019409011f71', 'Meteor Shower', 'https://l2db.ru/themes/l2db/images/items/weapon_meteor_shower_i01.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "blant", "attack": 213}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('c466a9bc-bc76-49df-a72f-768d75202f11', 'Arcana Mace', 'https://l2db.ru/themes/l2db/images/items/weapon_arcana_mace_i01_0.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "blant", "attack": 225}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('8d640670-d4e8-4b27-b6c1-7e1cfc8748f5', 'Draconic Bow', 'https://l2db.ru/themes/l2db/images/items/weapon_draconic_bow_i00_0.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "bow", "attack": 581}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('df744b54-a4ea-4113-964a-f56acea51d63', 'Heaven''s Divider', 'https://l2db.ru/themes/l2db/images/items/weapon_heavens_divider_i00_0.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "sword", "attack": 342}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('c4939841-2f6e-4866-8b85-e472fa92c5fe', 'Dragon Slayer', 'https://l2db.ru/themes/l2db/images/items/weapon_dragon_slayer_i01.png', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932', null, '{"type": "sword", "attack": 282}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('5de914cf-952a-4f63-a83e-dc8359ba48e3', 'Keepcrawler''s Gutripper', 'https://wow.zamimg.com/modelviewer/live/webthumbs/item/168/183720.webp', 'b142a159-5cec-4d3b-a3d8-3b23de2564d8', null, '{"type": "dagger"}');
INSERT INTO public.item (id, name, image_url, game_id, description, data) VALUES ('c20827d4-068d-40fd-8a02-560b8041c457', 'Sunblade', 'https://wow.zamimg.com/uploads/screenshots/small/430489.jpg', 'b142a159-5cec-4d3b-a3d8-3b23de2564d8', null, '{"type": "dagger"}');

INSERT INTO public.inventory (id, user_id, game_id) VALUES ('caa00108-bca4-4acf-ae1c-c4c6bf239276', '30eff35d-46c2-4483-874f-3607e269447e', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('4954ea64-dd45-42cf-9226-55782bab4b64', '30eff35d-46c2-4483-874f-3607e269447e', 'b142a159-5cec-4d3b-a3d8-3b23de2564d8');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('8cd38cea-c6b7-4491-9175-e14c8ec885bf', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', '923b1c32-e00d-415e-9f9e-f83a2b07d52e');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('9c89cba7-46cc-498c-aead-3cffe21a02de', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', 'fdd969ec-d9ec-43e4-8c1c-c94b7e9dc932');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('f9bf12db-f9e2-47ee-9b91-503a0404994b', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', 'b142a159-5cec-4d3b-a3d8-3b23de2564d8');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('153de3ee-76b8-4a87-b001-0df4fbbdfeb8', 'ce0d8279-08cd-438a-9949-a740aca9592f', '923b1c32-e00d-415e-9f9e-f83a2b07d52e');
INSERT INTO public.inventory (id, user_id, game_id) VALUES ('932057cc-f9d1-47a8-9e33-e2190cfb6c5e', 'ce0d8279-08cd-438a-9949-a740aca9592f', 'b142a159-5cec-4d3b-a3d8-3b23de2564d8');

INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('caa00108-bca4-4acf-ae1c-c4c6bf239276', '046b5be3-5aef-4407-8bf5-019409011f71', 2);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('caa00108-bca4-4acf-ae1c-c4c6bf239276', '8d640670-d4e8-4b27-b6c1-7e1cfc8748f5', 1);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('f9bf12db-f9e2-47ee-9b91-503a0404994b', '5de914cf-952a-4f63-a83e-dc8359ba48e3', 2);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('9c89cba7-46cc-498c-aead-3cffe21a02de', '046b5be3-5aef-4407-8bf5-019409011f71', 1);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('9c89cba7-46cc-498c-aead-3cffe21a02de', '8d640670-d4e8-4b27-b6c1-7e1cfc8748f5', 3);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('932057cc-f9d1-47a8-9e33-e2190cfb6c5e', 'c20827d4-068d-40fd-8a02-560b8041c457', 1);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('caa00108-bca4-4acf-ae1c-c4c6bf239276', 'ccdf8f74-8d3d-42d2-bd13-c98c9ef34834', 2);
INSERT INTO public.inventory_item (inventory_id, item_id, count) VALUES ('caa00108-bca4-4acf-ae1c-c4c6bf239276', 'c466a9bc-bc76-49df-a72f-768d75202f11', 6);

INSERT INTO public.lot (id, user_id, begin_price, buy_price, created_at, end_at, is_closed, max_bet_price) VALUES ('e4abbf77-0359-49ae-a945-03d0f2e573fb', 'ce0d8279-08cd-438a-9949-a740aca9592f', 200, 5000, '2020-12-10 00:00:00.000000', '2020-12-11 00:00:00.000000', 1, 500);
INSERT INTO public.lot (id, user_id, begin_price, buy_price, created_at, end_at, is_closed, max_bet_price) VALUES ('aa9efb6a-23ae-48c6-b17e-f97106d92f98', 'ce0d8279-08cd-438a-9949-a740aca9592f', 500, 1000, '2020-12-12 00:00:00.000000', '2020-12-19 00:00:00.000000', 0, 1000);
INSERT INTO public.lot (id, user_id, begin_price, buy_price, created_at, end_at, is_closed, max_bet_price) VALUES ('e3aba2d2-47c9-440e-8e5b-cf753408a06b', '30eff35d-46c2-4483-874f-3607e269447e', 100, 15000, '2020-12-12 00:00:00.000000', '2020-12-26 00:00:00.000000', 0, 6000);

INSERT INTO public.lot_item (lot_id, item_id, count) VALUES ('e3aba2d2-47c9-440e-8e5b-cf753408a06b', 'ccdf8f74-8d3d-42d2-bd13-c98c9ef34834', 1);
INSERT INTO public.lot_item (lot_id, item_id, count) VALUES ('e4abbf77-0359-49ae-a945-03d0f2e573fb', '8d640670-d4e8-4b27-b6c1-7e1cfc8748f5', 1);
INSERT INTO public.lot_item (lot_id, item_id, count) VALUES ('aa9efb6a-23ae-48c6-b17e-f97106d92f98', '8d640670-d4e8-4b27-b6c1-7e1cfc8748f5', 1);
INSERT INTO public.lot_item (lot_id, item_id, count) VALUES ('e3aba2d2-47c9-440e-8e5b-cf753408a06b', '046b5be3-5aef-4407-8bf5-019409011f71', 1);

INSERT INTO public.bet (id, lot_id, user_id, total_price, created_at) VALUES ('60f9671d-3ac8-4bc7-91c3-0e95186a77bf', 'e4abbf77-0359-49ae-a945-03d0f2e573fb', '30eff35d-46c2-4483-874f-3607e269447e', 500, '2020-12-12 00:00:00.000000');
INSERT INTO public.bet (id, lot_id, user_id, total_price, created_at) VALUES ('a4779dbe-1e7f-47df-85d8-83babdaf5d0d', 'e3aba2d2-47c9-440e-8e5b-cf753408a06b', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', 3000, '2020-12-12 00:00:00.000000');
INSERT INTO public.bet (id, lot_id, user_id, total_price, created_at) VALUES ('f8e1c69a-67c6-4462-ae48-5988b99055b0', 'e3aba2d2-47c9-440e-8e5b-cf753408a06b', 'ce0d8279-08cd-438a-9949-a740aca9592f', 5000, '2020-12-12 00:00:00.000000');
INSERT INTO public.bet (id, lot_id, user_id, total_price, created_at) VALUES ('fca2e632-41f9-483f-98b2-f8f1fa7736db', 'aa9efb6a-23ae-48c6-b17e-f97106d92f98', '30eff35d-46c2-4483-874f-3607e269447e', 1000, '2020-12-12 00:00:00.000000');
INSERT INTO public.bet (id, lot_id, user_id, total_price, created_at) VALUES ('2ed1915b-fab1-4ccf-b57c-cd993c9828e9', 'e3aba2d2-47c9-440e-8e5b-cf753408a06b', 'e219240d-7b25-4f4a-8b98-c4f90f9b6790', 6000, '2020-12-12 00:00:00.000000');

