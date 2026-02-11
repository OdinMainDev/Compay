CREATE TABLE "log_transactions" (
    "id" bigserial,
    "payment_id" bigint,
    "payment_trx_id" varchar(36),
    "withdraw_id" bigint,
    "withdraw_trx_id" varchar(36),
    "merchant_wallet_id" bigint,
    "merchant_id" int,
    "currency" varchar(20),
    "is_complete" boolean,
    "is_spent" boolean,
    "start_block" bigint,
    "tx_id" varchar(255),
    "tx_hash" varchar(255),
    "address" varchar(128),
    "data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_log_transactions_payment_id ON "log_transactions" ("payment_id");
CREATE INDEX idx_log_transactions_withdraw_id ON "log_transactions" ("withdraw_id");
CREATE INDEX idx_log_transactions_merchant_wallet_id ON "log_transactions" ("merchant_wallet_id");
CREATE INDEX idx_log_transactions_merchant_id ON "log_transactions" ("merchant_id");
CREATE INDEX idx_log_transactions_tx_id ON "log_transactions" ("tx_id");
CREATE UNIQUE INDEX idx_log_transactions_payment_id_tx_id ON "log_transactions" ("payment_id", "tx_id");
CREATE INDEX idx_log_transactions_wallet_complete_spent ON "log_transactions" ("merchant_wallet_id", "is_complete", "is_spent");
CREATE INDEX idx_log_transactions_address ON "log_transactions" ("address");