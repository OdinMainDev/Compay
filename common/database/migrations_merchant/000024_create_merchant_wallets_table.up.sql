CREATE TABLE "merchant_wallets" (
    "id" bigserial,
    "merchant_id" int,
    "payment_id" bigint,
    "currency" varchar(20),
    "address" varchar(128),
    "amount" int,
    "received_amount" int,
    "sent_amount" int,
    "private_key" text,
    "private_key_wif" text,
    "public_key" text,
    "public_key_full" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);