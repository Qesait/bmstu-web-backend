-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

\connect "container_logistics_6";

DROP TABLE IF EXISTS "containers";
CREATE TABLE "public"."containers" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "marking" character varying(11) NOT NULL,
    "type" character varying(50) NOT NULL,
    "length" bigint NOT NULL,
    "height" bigint NOT NULL,
    "width" bigint NOT NULL,
    "image_url" character varying(100),
    "is_deleted" boolean DEFAULT false NOT NULL,
    "cargo" character varying(50) NOT NULL,
    "weight" bigint NOT NULL,
    CONSTRAINT "containers_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "containers" ("uuid", "marking", "type", "length", "height", "width", "image_url", "is_deleted", "cargo", "weight") VALUES
('a20163ce-7be5-46ec-a50f-a313476b2bd1',	'CCCU6543210',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	'localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.jpg',	'f',	'Зерно',	15000),
('0706419e-b024-469d-a354-9480cd79c6a5',	'AAAU1234560',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	'localhost:9000/images/0706419e-b024-469d-a354-9480cd79c6a5.jpeg',	'f',	'Автомобиль',	5000),
('8f157a95-dad1-43e0-9372-93b51de06163',	'DDDU6543210',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	'localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.jpg',	'f',	'Фрукты',	13000),
('07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2',	'BBBU6543210',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	'localhost:9000/images/07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2.jpg',	'f',	'Телевизоры',	19000);

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "transportation_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "container_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("transportation_id", "container_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("transportation_id", "container_id") VALUES
('b0247ccd-28ab-45be-9680-f24213cf7aab',	'07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2');

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
('b0247ccd-28ab-45be-9680-f24213cf7aab',	'удалён',	'2023-10-25',	'2023-10-25',	'2023-10-25',	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'корабль');

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

ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_container" FOREIGN KEY (container_id) REFERENCES containers(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_transportation" FOREIGN KEY (transportation_id) REFERENCES transportations(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_customer" FOREIGN KEY (customer_id) REFERENCES users(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_moderator" FOREIGN KEY (moderator_id) REFERENCES users(uuid) NOT DEFERRABLE;

-- 2023-10-24 23:28:08.033597+00
