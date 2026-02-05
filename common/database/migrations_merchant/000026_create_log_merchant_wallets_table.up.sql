CREATE TABLE "log_merchant_wallets" (
    "id" bigserial,
    "merchant_id" int,
    "payment_id" bigint,
    "withdraw_id" bigint,
    "payment_trx_id" varchar(36),
    "withdraw_trx_id" varchar(36),
    "currency" varchar(20),
    "amount" int,
    "amount_usd" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);