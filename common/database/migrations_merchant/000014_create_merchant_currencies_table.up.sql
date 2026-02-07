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