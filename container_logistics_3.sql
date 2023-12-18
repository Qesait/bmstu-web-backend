-- Adminer 4.8.1 PostgreSQL 16.1 (Debian 16.1-1.pgdg120+1) dump

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
('8f157a95-dad1-43e0-9372-93b51de06163',	'DDDU6543210',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	'localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.jpg',	'f',	'Фрукты',	13000),
('07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2',	'BBBU6543210',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	'localhost:9000/images/07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2.jpg',	'f',	'Телевизоры',	19000),
('87f1e44c-2d5f-4611-ae7c-4b56ba690d05',	'XYZU9876543',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	'localhost:9000/images/87f1e44c-2d5f-4611-ae7c-4b56ba690d05.jpg',	'f',	'Медицинское оборудование',	8000),
('f4c76108-9cd1-42e8-81d5-803eaed7866a',	'JRZU1176543',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	'localhost:9000/images/f4c76108-9cd1-42e8-81d5-803eaed7866a.jpg',	'f',	'Одежда',	16000),
('0706419e-b024-469d-a354-9480cd79c6a5',	'AAAU1234560',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	NULL,	'f',	'Автомобиль',	5000),
('a20163ce-7be5-46ec-a50f-a313476b2bd1',	'CCCU6543210',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	'localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.jpg',	'f',	'Зерно',	15000);

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "transportation_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "container_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("transportation_id", "container_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("transportation_id", "container_id") VALUES
('ff48acb4-c78a-4561-b9d0-ba6c843cc5bc',	'a20163ce-7be5-46ec-a50f-a313476b2bd1');

DROP TABLE IF EXISTS "transportations";
CREATE TABLE "public"."transportations" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) NOT NULL,
    "creation_date" timestamp NOT NULL,
    "formation_date" timestamp,
    "completion_date" timestamp,
    "moderator_id" uuid,
    "customer_id" uuid NOT NULL,
    "transport" character varying(50),
    "delivery_status" character varying(40),
    CONSTRAINT "transportations_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "transportations" ("uuid", "status", "creation_date", "formation_date", "completion_date", "moderator_id", "customer_id", "transport", "delivery_status") VALUES
('ff48acb4-c78a-4561-b9d0-ba6c843cc5bc',	'сформирован',	'2023-12-10 18:59:10.519535',	NULL,	NULL,	NULL,	'ddc3bdd4-c79a-46cf-8558-064b4ba8f87c',	NULL,	'отменена');

DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "role" bigint,
    "login" character varying(30) NOT NULL,
    "password" character varying(40) NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "users" ("uuid", "role", "login", "password") VALUES
('e726583b-232a-4db6-b215-f39f4358b1d4',	1,	'regular user',	'40bd001563085fc35165329ea1ff5c5ecbdbbeef'),
('7af41851-c0cd-47d6-b2b0-a617a935221e',	2,	'admin',	'd033e22ae348aeb5660fc2140aec35850c4da997'),
('ddc3bdd4-c79a-46cf-8558-064b4ba8f87c',	1,	'user',	'2ddbd6d791999fa766ed57a9831fa554fac2f6ae');

ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_container" FOREIGN KEY (container_id) REFERENCES containers(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_transportation" FOREIGN KEY (transportation_id) REFERENCES transportations(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_customer" FOREIGN KEY (customer_id) REFERENCES users(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_moderator" FOREIGN KEY (moderator_id) REFERENCES users(uuid) NOT DEFERRABLE;

-- 2023-12-18 21:58:09.165759+00
