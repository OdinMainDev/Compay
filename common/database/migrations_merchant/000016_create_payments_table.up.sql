CREATE TABLE "payments" (
    "id" bigserial,
    "trx_id" varchar(36),
    "merchant_id" int,
    "merchant_wallet_id" bigint,
    "currency" varchar(20),
    "status" varchar(20),
    "address" varchar(128),
    "amount" int,
    "amount_usd" int,
    "received_amount" int,
    "received_amount_usd" int,
    "required_confirmations" int,
    "expires_at" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);