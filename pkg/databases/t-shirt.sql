CREATE TYPE "gender" AS ENUM (
  'male',
  'women',
  'n/a'
);

CREATE TYPE "size" AS ENUM (
  'xs',
  's',
  'm',
  'l',
  'xl'
);

CREATE TYPE "role" AS ENUM (
  'geust',
  'admin'
);

CREATE TABLE "products" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "gender" gender,
  "style" varchar(255),
  "size" size,
  "price" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_datetime" timestamp DEFAULT (now())
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "username" varchar,
  "gender" gender,
  "role" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_datetime" timestamp DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "user_id" integer,
  "status" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_datetime" timestamp DEFAULT (now()),
  "snap_order" jsonb
);

CREATE TABLE "shippings" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "order_id" integer,
  "address" varchar(255),
  "subdistricet" varchar(255),
  "district" varchar(255),
  "provind" varchar(255),
  "zip_code" varchar(255),
  "created_at" timestamp DEFAULT (now()),
  "updated_datetime" timestamp DEFAULT (now())
);

CREATE TABLE "order_product" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "order_id" integer,
  "product_id" integer,
  "created_at" timestamp DEFAULT (now()),
  "updated_datetime" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "products"."updated_datetime" IS 'ON UPDATE CURRENT_TIMESTAMP';

COMMENT ON COLUMN "users"."updated_datetime" IS 'ON UPDATE CURRENT_TIMESTAMP';

COMMENT ON COLUMN "orders"."updated_datetime" IS 'ON UPDATE CURRENT_TIMESTAMP';

COMMENT ON COLUMN "shippings"."updated_datetime" IS 'ON UPDATE CURRENT_TIMESTAMP';

COMMENT ON COLUMN "order_product"."updated_datetime" IS 'ON UPDATE CURRENT_TIMESTAMP';

ALTER TABLE "orders" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_product" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_product" ADD FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "shippings" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");
