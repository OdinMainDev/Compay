CREATE TABLE "merchant_wallets" (
    "id" bigserial,
    "merchant_id" int,
    "payment_id" bigint,
    "currency" varchar(20),
    "address" varchar(128),
    "amount" bigint,
    "received_amount" bigint,
    "sent_amount" bigint,
    "private_key" text,
    "private_key_wif" text,
    "public_key" text,
    "public_key_full" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_merchant_wallets_merchant_id ON "merchant_wallets" ("merchant_id");
CREATE INDEX idx_merchant_wallets_payment_id ON "merchant_wallets" ("payment_id");
CREATE INDEX idx_merchant_wallets_merchant_id_currency ON "merchant_wallets" ("merchant_id", "currency");
CREATE INDEX idx_merchant_wallets_address ON "merchant_wallets" ("address");