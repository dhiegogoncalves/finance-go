CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "type" varchar NOT NULL,
  "description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "categories" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" bigint NOT NULL,
  "category_id" bigint NOT NULL,
  "title" varchar NOT NULL,
  "type" varchar NOT NULL,
  "description" varchar NOT NULL,
  "value" decimal NOT NULL,
  "date" date NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "accounts" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");