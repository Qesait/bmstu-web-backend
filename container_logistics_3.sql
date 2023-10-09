DROP TABLE IF EXISTS "container_types";
DROP SEQUENCE IF EXISTS container_types_container_type_id_seq;
CREATE SEQUENCE container_types_container_type_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."container_types" (
    "container_type_id" bigint DEFAULT nextval('container_types_container_type_id_seq') NOT NULL,
    "name" character varying(50) NOT NULL,
    "length" bigint NOT NULL,
    "height" bigint NOT NULL,
    "width" bigint NOT NULL,
    "max_gross" bigint NOT NULL,
    CONSTRAINT "container_types_pkey" PRIMARY KEY ("container_type_id")
) WITH (oids = false);

INSERT INTO "container_types" ("container_type_id", "name", "length", "height", "width", "max_gross") VALUES
(4,	'40 футовый контейнер увеличенной высоты	',	12045,	2596,	2350,	26700),
(3,	'Стандартный 40-ка футовый контейнер',	12045,	2381,	2350,	26700),
(2,	'20 футовый контейнер увеличенной высоты',	5905,	2596,	2350,	21650),
(1,	'Стандартный 20-ти футовый контейнер',	5905,	2381,	2350,	21770);

DROP TABLE IF EXISTS "containers";
DROP SEQUENCE IF EXISTS containers_container_id_seq;
CREATE SEQUENCE containers_container_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."containers" (
    "container_id" bigint DEFAULT nextval('containers_container_id_seq') NOT NULL,
    "type_id" bigint NOT NULL,
    "image_url" character varying(100) NOT NULL,
    "is_deleted" boolean NOT NULL,
    "purchase_date" date NOT NULL,
    "cargo" character varying(50) NOT NULL,
    "weight" bigint NOT NULL,
    "marking" character varying(11) NOT NULL,
    CONSTRAINT "containers_pkey" PRIMARY KEY ("container_id")
) WITH (oids = false);

INSERT INTO "containers" ("container_id", "type_id", "image_url", "is_deleted", "purchase_date", "cargo", "weight", "marking") VALUES
(2,	3,	'1.jpg',	'f',	'2020-12-23',	'Телевизоры',	19000,	'BBBU6543210'),
(1,	1,	'0.jpeg',	't',	'2020-08-09',	'',	0,	'AAAU1234560'),
(4,	2,	'4.jpg',	'f',	'2019-04-14',	'Фрукты',	13000,	'DDDU6543210'),
(3,	1,	'3.jpg',	'f',	'2021-07-19',	'Зерно',	15000,	'CCCU6543210');

DROP TABLE IF EXISTS "statuses";
DROP SEQUENCE IF EXISTS statuses_status_id_seq;
CREATE SEQUENCE statuses_status_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."statuses" (
    "status_id" bigint DEFAULT nextval('statuses_status_id_seq') NOT NULL,
    "name" character varying(50) NOT NULL,
    CONSTRAINT "statuses_pkey" PRIMARY KEY ("status_id")
) WITH (oids = false);

INSERT INTO "statuses" ("status_id", "name") VALUES
(1,	'введён'),
(2,	'в работе'),
(3,	'завершён'),
(4,	'отменён'),
(5,	'удалён');

DROP TABLE IF EXISTS "transportation_compositions";
CREATE TABLE "public"."transportation_compositions" (
    "container_id" bigint NOT NULL,
    "transportation_id" bigint NOT NULL,
    CONSTRAINT "transportation_compositions_pkey" PRIMARY KEY ("container_id", "transportation_id")
) WITH (oids = false);

INSERT INTO "transportation_compositions" ("container_id", "transportation_id") VALUES
(2,	2),
(3,	3),
(4,	4);

DROP TABLE IF EXISTS "transportations";
DROP SEQUENCE IF EXISTS transportations_transportation_id_seq;
CREATE SEQUENCE transportations_transportation_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."transportations" (
    "transportation_id" bigint DEFAULT nextval('transportations_transportation_id_seq') NOT NULL,
    "status_id" bigint NOT NULL,
    "creation_date" date NOT NULL,
    "formation_date" date,
    "completion_date" date,
    "moderator_id" bigint NOT NULL,
    "customer_id" bigint NOT NULL,
    "transport_vehicle" character varying(50) NOT NULL,
    CONSTRAINT "transportations_pkey" PRIMARY KEY ("transportation_id")
) WITH (oids = false);

INSERT INTO "transportations" ("transportation_id", "status_id", "creation_date", "formation_date", "completion_date", "moderator_id", "customer_id", "transport_vehicle") VALUES
(1,	1,	'2023-09-26',	NULL,	NULL,	2,	1,	'Корабль'),
(2,	2,	'2023-08-12',	'2023-09-14',	NULL,	2,	1,	'Поезд'),
(3,	3,	'2023-07-10',	'2023-07-14',	'2023-08-09',	2,	1,	'Корабль'),
(4,	4,	'2023-09-25',	NULL,	NULL,	2,	1,	'Поезд');

DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_user_id_seq;
CREATE SEQUENCE users_user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."users" (
    "user_id" bigint DEFAULT nextval('users_user_id_seq') NOT NULL,
    "login" character varying(30) NOT NULL,
    "password" character varying(30) NOT NULL,
    "name" character varying(50) NOT NULL,
    "moderator" boolean NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("user_id")
) WITH (oids = false);

INSERT INTO "users" ("user_id", "login", "password", "name", "moderator") VALUES
(1,	'user 1',	'password 1',	'Просто пользователь',	'f'),
(2,	'user 2',	'password 2',	'Модератор',	't');

ALTER TABLE ONLY "public"."containers" ADD CONSTRAINT "fk_containers_container_type" FOREIGN KEY (type_id) REFERENCES container_types(container_type_id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_container" FOREIGN KEY (container_id) REFERENCES containers(container_id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportation_compositions" ADD CONSTRAINT "fk_transportation_compositions_transportation" FOREIGN KEY (transportation_id) REFERENCES transportations(transportation_id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_customer" FOREIGN KEY (customer_id) REFERENCES users(user_id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_moderator" FOREIGN KEY (moderator_id) REFERENCES users(user_id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."transportations" ADD CONSTRAINT "fk_transportations_status" FOREIGN KEY (status_id) REFERENCES statuses(status_id) NOT DEFERRABLE;