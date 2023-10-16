-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

\connect "container_logistics_4";

DROP TABLE IF EXISTS "container_types";
CREATE TABLE "public"."container_types" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "name" character varying(50) NOT NULL,
    "length" bigint NOT NULL,
    "height" bigint NOT NULL,
    "width" bigint NOT NULL,
    "max_gross" bigint NOT NULL,
    CONSTRAINT "container_types_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "container_types" ("uuid", "name", "length", "height", "width", "max_gross") VALUES
('29872d1b-ad47-4c6c-a690-c784c4d1d0ac',	'40 футовый контейнер увеличенной высоты	',	12045,	2596,	2350,	26700),
('6c04dd7d-38c9-4f04-ac04-d5ec18bc87e0',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	26700),
('0884d0b9-a164-42f4-8f65-5ea16759c63a',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	21650),
('ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	21770);

DROP TABLE IF EXISTS "containers";
CREATE TABLE "public"."containers" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "type_id" uuid NOT NULL,
    "image_url" character varying(100) NOT NULL,
    "is_deleted" boolean NOT NULL,
    "purchase_date" date NOT NULL,
    "cargo" character varying(50) NOT NULL,
    "weight" bigint NOT NULL,
    "marking" character varying(11) NOT NULL,
    CONSTRAINT "containers_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "containers" ("uuid", "type_id", "image_url", "is_deleted", "purchase_date", "cargo", "weight", "marking") VALUES
('34625fe0-39ab-4297-8198-cf707e6e7e02',	'6c04dd7d-38c9-4f04-ac04-d5ec18bc87e0',	'1.jpg',	'f',	'2020-12-23',	'Телевизоры',	19000,	'BBBU6543210'),
('7f8f5915-2694-4fe9-b887-1aa3a259625e',	'ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'0.jpeg',	't',	'2020-08-09',	'',	0,	'AAAU1234560'),
('f9d02ce0-821c-47ca-93eb-b5b2ffa8f2de',	'ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'3.jpg',	'f',	'2021-07-19',	'Зерно',	15000,	'CCCU6543210'),
('37784fff-414f-43f8-802f-cd5eea168f3b',	'0884d0b9-a164-42f4-8f65-5ea16759c63a',	'4.jpg',	'f',	'2019-04-14',	'Фрукты',	13000,	'DDDU6543210');

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "transportation_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "container_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("transportation_id", "container_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("transportation_id", "container_id") VALUES
('1d4c768c-2c5a-41c9-8062-7b348802fbb0',	'34625fe0-39ab-4297-8198-cf707e6e7e02'),
('9e044142-e94e-4dd4-80e7-38d73d397d02',	'f9d02ce0-821c-47ca-93eb-b5b2ffa8f2de'),
('35ddf42f-8a34-464d-8c32-a470be07ea73',	'37784fff-414f-43f8-802f-cd5eea168f3b'),
('7a31f7e3-6a23-41c5-b11e-4dfa6eb730f9',	'34625fe0-39ab-4297-8198-cf707e6e7e02'),
('7a31f7e3-6a23-41c5-b11e-4dfa6eb730f9',	'7f8f5915-2694-4fe9-b887-1aa3a259625e'),
('225c7f60-228c-4dfb-ae16-5c234c3a0d3e',	'37784fff-414f-43f8-802f-cd5eea168f3b'),
('225c7f60-228c-4dfb-ae16-5c234c3a0d3e',	'f9d02ce0-821c-47ca-93eb-b5b2ffa8f2de'),
('c864db1c-9192-4bb2-a5ae-7717b9070c6c',	'34625fe0-39ab-4297-8198-cf707e6e7e02');

DROP TABLE IF EXISTS "transportations";
CREATE TABLE "public"."transportations" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) DEFAULT 'введён' NOT NULL,
    "creation_date" date NOT NULL,
    "formation_date" date,
    "completion_date" date,
    "moderator_id" uuid,
    "customer_id" uuid,
    "transport_vehicle" character varying(50) NOT NULL,
    CONSTRAINT "transportations_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "transportations" ("uuid", "status", "creation_date", "formation_date", "completion_date", "moderator_id", "customer_id", "transport_vehicle") VALUES
('1d4c768c-2c5a-41c9-8062-7b348802fbb0',	'в работе',	'2023-08-12',	'2023-09-14',	NULL,	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'Поезд'),
('9e044142-e94e-4dd4-80e7-38d73d397d02',	'завершён',	'2023-07-10',	'2023-07-14',	'2023-08-09',	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'Корабль'),
('35ddf42f-8a34-464d-8c32-a470be07ea73',	'отменён',	'2023-09-25',	NULL,	NULL,	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'Поезд'),
('7a31f7e3-6a23-41c5-b11e-4dfa6eb730f9',	'в работе',	'2023-09-26',	NULL,	NULL,	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'Корабль'),
('225c7f60-228c-4dfb-ae16-5c234c3a0d3e',	'удалён',	'2023-10-12',	NULL,	NULL,	NULL,	NULL,	'машина'),
('c864db1c-9192-4bb2-a5ae-7717b9070c6c',	'удалён',	'2023-10-16',	NULL,	NULL,	NULL,	NULL,	'машина');

DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "login" character varying(30) NOT NULL,
    "password" character varying(30) NOT NULL,
    "name" character varying(50) NOT NULL,
    "moderator" boolean NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "users" ("uuid", "login", "password", "name", "moderator") VALUES
('5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'user 1',	'password 1',	'Просто пользователь',	'f'),
('796c70e1-5f27-4433-a415-95e7272effa5',	'user 2',	'password 2',	'Модератор',	't');

ALTER TABLE ONLY "public"."containers" ADD CONSTRAINT "fk_containers_container_type" FOREIGN KEY (type_id) REFERENCES container_types(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_container" FOREIGN KEY (container_id) REFERENCES containers(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_transportation" FOREIGN KEY (transportation_id) REFERENCES transportations(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_customer" FOREIGN KEY (customer_id) REFERENCES users(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_moderator" FOREIGN KEY (moderator_id) REFERENCES users(uuid) NOT DEFERRABLE;

-- 2023-10-16 13:09:50.561342+00
