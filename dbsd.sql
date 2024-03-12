CREATE TYPE "user_role" AS ENUM (
  'super',
  'admin'
);

CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "username" varchar UNIQUE,
  "password" varchar,
  "role" user_role,
  "active" bool DEFAULT false,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "workshop_orders" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "address" varchar,
  "dni" varchar,
  "ruc" varchar,
  "brand" varchar,
  "model" varchar,
  "color" varchar,
  "plate" varchar,
  "mileage" varchar,
  "observation" text,
  "created_at" timestamp DEFAULT (now()),
  "user" int
);

CREATE TABLE "workshop_order_items" (
  "id" int PRIMARY KEY,
  "code" varchar,
  "quantity" int,
  "price" float,
  "description" text,
  "created_at" timestamp DEFAULT (now()),
  "order" int
);

ALTER TABLE "workshop_orders" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

ALTER TABLE "workshop_order_items" ADD FOREIGN KEY ("order") REFERENCES "workshop_orders" ("id");
