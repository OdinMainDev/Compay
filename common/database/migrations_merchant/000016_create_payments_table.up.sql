CREATE TABLE "payments" (
    "id" bigserial,
    "trx_id" varchar(36),
    "merchant_id" int,
    "currency" varchar(20),
    "address" varchar(255),
    "amount" int,
    "amount_usd" int,
    "private_key" text,
    "private_key_wif" text,
    "public_key" text,
    "public_key_full" text,
    "is_active" boolean,
    "expires_at" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);