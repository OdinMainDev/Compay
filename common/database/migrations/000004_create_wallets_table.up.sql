CREATE TABLE "wallets" (
    "id" bigserial,
    "user_id" int,
    "address" varchar(255) NOT NULL,
    "network" varchar(255) NOT NULL,
    "private_key" text NOT NULL,
    "mnemonic" text NOT NULL,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_wallets_user_id ON "wallets" ("user_id");
CREATE INDEX idx_wallets_address ON "wallets" ("address");
CREATE INDEX idx_wallets_network ON "wallets" ("network");
