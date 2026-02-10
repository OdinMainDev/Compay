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

CREATE INDEX idx_merchant_stats_merchant_id ON "merchant_stats" ("merchant_id");
CREATE UNIQUE INDEX idx_merchant_stats_merchant_id_currency ON "merchant_stats" ("merchant_id", "currency");