CREATE TABLE "payments" (
    "id" bigserial,
    "trx_id" varchar(36),
    "merchant_id" int,
    "merchant_wallet_id" bigint,
    "currency" varchar(20),
    "status" varchar(60),
    "address" varchar(128),
    "amount" bigint,
    "amount_usd" bigint,
    "received_amount" bigint,
    "received_amount_usd" bigint,
    "required_confirmations" bigint,
    "expires_at" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX idx_payments_trx_id ON "payments" ("trx_id");
CREATE INDEX idx_payments_merchant_id ON "payments" ("merchant_id");
CREATE INDEX idx_payments_merchant_wallet_id ON "payments" ("merchant_wallet_id");
CREATE INDEX idx_payments_merchant_id_status ON "payments" ("merchant_id", "status");
CREATE INDEX idx_payments_currency ON "payments" ("currency");
CREATE INDEX idx_payments_address ON "payments" ("address");
CREATE INDEX idx_payments_status ON "payments" ("status");
CREATE INDEX idx_payments_created_at ON "payments" ("created_at" DESC);