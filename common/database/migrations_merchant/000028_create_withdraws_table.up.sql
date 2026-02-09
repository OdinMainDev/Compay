CREATE TABLE "withdraws" (
    "id" bigserial,
    "trx_id" varchar(36),
    "merchant_id" int,
    "merchant_wallet_id" int,
    "currency" varchar(20),
    "is_complete" boolean,
    "address" varchar(128),
    "amount" int,
    "amount_usd" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);