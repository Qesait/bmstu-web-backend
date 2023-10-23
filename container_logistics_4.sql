-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

\connect "container_logistics_5";

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
('ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	21770),
('0884d0b9-a164-42f4-8f65-5ea16759c63a',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	21650);

DROP TABLE IF EXISTS "containers";
CREATE TABLE "public"."containers" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "type_id" uuid NOT NULL,
    "image_url" character varying(100) NOT NULL,
    "is_deleted" boolean DEFAULT false NOT NULL,
    "purchase_date" date NOT NULL,
    "cargo" character varying(50) NOT NULL,
    "weight" bigint NOT NULL,
    "marking" character varying(11) NOT NULL,
    CONSTRAINT "containers_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "containers" ("uuid", "type_id", "image_url", "is_deleted", "purchase_date", "cargo", "weight", "marking") VALUES
('34625fe0-39ab-4297-8198-cf707e6e7e02',	'6c04dd7d-38c9-4f04-ac04-d5ec18bc87e0',	'1.jpg',	'f',	'2020-12-23',	'Телевизоры',	19000,	'BBBU6543210'),
('f9d02ce0-821c-47ca-93eb-b5b2ffa8f2de',	'ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'3.jpg',	'f',	'2021-07-19',	'Зерно',	15000,	'CCCU6543210'),
('37784fff-414f-43f8-802f-cd5eea168f3b',	'0884d0b9-a164-42f4-8f65-5ea16759c63a',	'4.jpg',	'f',	'2019-04-14',	'Фрукты',	13000,	'DDDU6543210'),
('26fa0660-eaf2-466c-9efe-e757cf32680e',	'29872d1b-ad47-4c6c-a690-c784c4d1d0ac',	'1.jpg',	't',	'2021-07-19',	'какие-то товары',	10001,	'EEEU3539010'),
('045c6172-1149-414a-8264-87edf10997e7',	'6c04dd7d-38c9-4f04-ac04-d5ec18bc87e0',	'0.jpeg',	't',	'2021-07-19',	'какие-то товары',	10000,	'EEEU3539010'),
('7f8f5915-2694-4fe9-b887-1aa3a259625e',	'ba308d64-1c19-4c60-a7b3-e53127ad69c4',	'0.jpeg',	'f',	'2020-08-09',	'',	0,	'AAAU1234560');

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "transportation_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "container_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("transportation_id", "container_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("transportation_id", "container_id") VALUES
('be114237-b4ca-43ee-b12f-4dd0e395d253',	'7f8f5915-2694-4fe9-b887-1aa3a259625e'),
('37218013-3e9c-4376-988b-f6b466776559',	'34625fe0-39ab-4297-8198-cf707e6e7e02'),
('37218013-3e9c-4376-988b-f6b466776559',	'7f8f5915-2694-4fe9-b887-1aa3a259625e');

DROP TABLE IF EXISTS "transportations";
CREATE TABLE "public"."transportations" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) NOT NULL,
    "creation_date" date NOT NULL,
    "formation_date" date,
    "completion_date" date,
    "moderator_id" uuid,
    "customer_id" uuid NOT NULL,
    "transport" character varying(50) NOT NULL,
    CONSTRAINT "transportations_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "transportations" ("uuid", "status", "creation_date", "formation_date", "completion_date", "moderator_id", "customer_id", "transport") VALUES
('be114237-b4ca-43ee-b12f-4dd0e395d253',	'удалён',	'2023-10-23',	'2023-10-23',	'2023-10-23',	NULL,	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'корабль'),
('37218013-3e9c-4376-988b-f6b466776559',	'сформирован',	'2023-10-23',	'2023-10-23',	NULL,	NULL,	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'поезд');

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

-- 2023-10-23 18:18:03.074816+00