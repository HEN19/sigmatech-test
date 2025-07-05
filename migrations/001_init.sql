-- +migrate Up
CREATE TABLE "users" (
  "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
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


CREATE TABLE customers (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "nik" VARCHAR(255) UNIQUE NOT NULL,
    "full_name" VARCHAR(255) NOT NULL,
    "legal_name" VARCHAR(255) NOT NULL,
    "birth_place" VARCHAR(255) NOT NULL,
    "birth_date" VARCHAR(10) NOT NULL,
    "salary" DECIMAL(15,2) NOT NULL,
    "ktp_photo" TEXT NOT NULL,
    "selfie_photo" TEXT NOT NULL,
    "limit_1_month" DECIMAL(15,2) DEFAULT 0,
    "limit_2_month" DECIMAL(15,2) DEFAULT 0,
    "limit_3_month" DECIMAL(15,2) DEFAULT 0,
    "limit_6_month" DECIMAL(15,2) DEFAULT 0,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted" BOOLEAN DEFAULT FALSE

);

CREATE TABLE transactions (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "contract_number" VARCHAR(255) UNIQUE NOT NULL,
    "otr" DECIMAL(15,2) NOT NULL,
    "admin_fee" DECIMAL(15,2) NOT NULL,
    "installment" INT NOT NULL,
    "interest" DECIMAL(5,2) NOT NULL,
    "asset_name" VARCHAR(255) NOT NULL,
    "customer_id" UUID REFERENCES customers(id) ON DELETE CASCADE,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted" BOOLEAN DEFAULT FALSE
);

CREATE INDEX idx_transactions_customer_id ON transactions(customer_id);


-- +migrate StatementEnd