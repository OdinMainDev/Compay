CREATE TABLE "withdraws" (
    "id" bigserial,
    "trx_id" varchar(36),
    "merchant_id" int,
    "merchant_wallet_id" bigint,
    "currency" varchar(20),
    "is_complete" boolean,
    "address" varchar(128),
    "amount" bigint,
    "amount_usd" bigint,
    "error" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX idx_withdraws_trx_id ON "withdraws" ("trx_id");
CREATE INDEX idx_withdraws_merchant_id ON "withdraws" ("merchant_id");
CREATE INDEX idx_withdraws_merchant_wallet_id ON "withdraws" ("merchant_wallet_id");
CREATE INDEX idx_withdraws_wallet_is_complete ON "withdraws" ("merchant_wallet_id", "is_complete");
CREATE INDEX idx_withdraws_is_complete ON "withdraws" ("is_complete");