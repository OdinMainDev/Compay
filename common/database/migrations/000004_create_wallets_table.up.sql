CREATE TABLE "wallets" (
    "id" bigserial,
    "user_id" int,
    "mnemonic" varchar(320),
    "private_key" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_wallets_user_id ON "wallets" ("user_id");
