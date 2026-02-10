CREATE TABLE "log_merchant_wallets" (
    "id" bigserial,
    "merchant_id" int,
    "merchant_wallet_id" bigint,
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

CREATE INDEX idx_log_merchant_wallets_merchant_id ON "log_merchant_wallets" ("merchant_id");
CREATE INDEX idx_log_merchant_wallets_merchant_wallet_id ON "log_merchant_wallets" ("merchant_wallet_id");
CREATE INDEX idx_log_merchant_wallets_payment_id ON "log_merchant_wallets" ("payment_id");
CREATE INDEX idx_log_merchant_wallets_withdraw_id ON "log_merchant_wallets" ("withdraw_id");