CREATE TABLE "merchant_stats" (
    "id" bigserial,
    "merchant_id" int,
    "currency" varchar(20),
    "total_wallets" int,
    "total_transactions" int,
    "total_funds" bigint,
    "total_funds_usd" bigint,
    "avail_funds" bigint,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);