CREATE TABLE "wallet_address" (
    "id" bigserial,
    "wallet_id" int,
    "currency" varchar(20),
    "address" varchar(128),
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_wallet_address_currency ON "wallet_address" ("currency");
CREATE INDEX idx_wallet_address_wallet_id ON "wallet_address" ("wallet_id");

CREATE INDEX IF NOT EXISTS "idx_wallet_address_wallet_id_currency"
    ON "wallet_address" ("wallet_id", "currency");
