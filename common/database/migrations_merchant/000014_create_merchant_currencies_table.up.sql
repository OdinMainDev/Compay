CREATE TABLE "merchant_currencies" (
    "id" bigserial,
    "merchant_id" int,
    "currency" varchar(20),
    "creator_id" bigint,
    "min_amount" int,
    "max_amount" int,
    "min_amount_withdraw" int,
    "max_amount_withdraw" int,
    "confirms" int,
    "invoice_time" int,
    "estimated_fee_block_count" int,
    "fee_type" varchar(20),
    "fee_fixed" int,
    "fee_percent" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_merchant_currencies_merchant_id ON "merchant_currencies" ("merchant_id");
CREATE UNIQUE INDEX idx_merchant_currencies_merchant_id_currency ON "merchant_currencies" ("merchant_id", "currency");
CREATE INDEX idx_merchant_currencies_creator_id ON "merchant_currencies" ("creator_id");