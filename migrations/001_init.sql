-- +migrate Up
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "username" VARCHAR(100) NOT NULL,
  "password" VARCHAR(256) NOT NULL,
  "first_name" VARCHAR(75) NOT NULL,
  "last_name" VARCHAR(75) NOT NULL,
  "gender" CHAR(1) NOT NULL,
  "phone" VARCHAR(20) NOT NULL,
  "email" VARCHAR(100) NOT NULL,
  "address" TEXT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate StatementEnd