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
('a20163ce-7be5-46ec-a50f-a313476b2bd1',	'CCCU6543210',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	'localhost:9000/images/a20163ce-7be5-46ec-a50f-a313476b2bd1.jpg',	'f',	'Зерно',	15000),
('0706419e-b024-469d-a354-9480cd79c6a5',	'AAAU1234560',	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	'localhost:9000/images/0706419e-b024-469d-a354-9480cd79c6a5.jpeg',	'f',	'Автомобиль',	5000),
('8f157a95-dad1-43e0-9372-93b51de06163',	'DDDU6543210',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	'localhost:9000/images/8f157a95-dad1-43e0-9372-93b51de06163.jpg',	'f',	'Фрукты',	13000),
('07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2',	'BBBU6543210',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	'localhost:9000/images/07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2.jpg',	'f',	'Телевизоры',	19000),
('87f1e44c-2d5f-4611-ae7c-4b56ba690d05',	'XYZU9876543',	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	'localhost:9000/images/87f1e44c-2d5f-4611-ae7c-4b56ba690d05.jpg',	'f',	'Медицинское оборудование',	8000),
('f4c76108-9cd1-42e8-81d5-803eaed7866a',	'XYZU9876543',	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	'localhost:9000/images/f4c76108-9cd1-42e8-81d5-803eaed7866a.jpg',	'f',	'Одежда',	16000);

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "transportation_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "container_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("transportation_id", "container_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("transportation_id", "container_id") VALUES
('b0247ccd-28ab-45be-9680-f24213cf7aab',	'07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2'),
('7be39e57-784b-437a-a222-834f7995b1b5',	'a20163ce-7be5-46ec-a50f-a313476b2bd1'),
('7be39e57-784b-437a-a222-834f7995b1b5',	'0706419e-b024-469d-a354-9480cd79c6a5'),
('47e5468f-6426-4d0b-9866-85d2d0dbbee4',	'8f157a95-dad1-43e0-9372-93b51de06163'),
('47e5468f-6426-4d0b-9866-85d2d0dbbee4',	'07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2'),
('b8a71f96-48c2-436e-b587-29a9074e0b5c',	'07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2'),
('b8a71f96-48c2-436e-b587-29a9074e0b5c',	'f4c76108-9cd1-42e8-81d5-803eaed7866a'),
('b8a71f96-48c2-436e-b587-29a9074e0b5c',	'87f1e44c-2d5f-4611-ae7c-4b56ba690d05'),
('2aa39aff-e4c3-4f39-ab3d-2dde1989569b',	'07d0cbdc-8e0f-4308-a7aa-11976ee6e5b2'),
('2aa39aff-e4c3-4f39-ab3d-2dde1989569b',	'0706419e-b024-469d-a354-9480cd79c6a5');

DROP TABLE IF EXISTS "transportations";
CREATE TABLE "public"."transportations" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) NOT NULL,
    "creation_date" timestamp NOT NULL,
    "formation_date" timestamp,
    "completion_date" timestamp,
    "moderator_id" uuid,
    "customer_id" uuid NOT NULL,
    "transport" character varying(50) NOT NULL,
    CONSTRAINT "transportations_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "transportations" ("uuid", "status", "creation_date", "formation_date", "completion_date", "moderator_id", "customer_id", "transport") VALUES
('b0247ccd-28ab-45be-9680-f24213cf7aab',	'удалён',	'2023-10-25 00:00:00',	'2023-10-25 00:00:00',	'2023-10-25 00:00:00',	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'корабль'),
('47e5468f-6426-4d0b-9866-85d2d0dbbee4',	'завершён',	'2023-11-21 14:17:04.546543',	'2023-11-21 14:17:45.030415',	'2023-11-21 14:17:54.529754',	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	''),
('b8a71f96-48c2-436e-b587-29a9074e0b5c',	'сформирован',	'2023-11-21 14:18:24.047468',	'2023-11-21 20:59:30.569044',	NULL,	NULL,	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	''),
('2aa39aff-e4c3-4f39-ab3d-2dde1989569b',	'черновик',	'2023-11-21 20:59:57.808767',	NULL,	NULL,	NULL,	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	''),
('7be39e57-784b-437a-a222-834f7995b1b5',	'отклонён',	'2023-11-21 14:15:44.663215',	'2023-11-21 14:16:53.356712',	NULL,	'796c70e1-5f27-4433-a415-95e7272effa5',	'5f58c307-a3f2-4b13-b888-c80ad08d5ed3',	'');

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

-- 2023-11-27 15:47:21.42604+00
