CREATE TABLE "merchant_currencies" (
    "id" bigserial,
    "merchant_id" int,
    "currency" varchar(20),
    "min_amount" int,
    "max_amount" int,
    "min_amount_withdraw" int,
    "max_amount_withdraw" int,
    "min_confirms" int,
    "fee_type" varchar(20),
    "fee_fixed" int,
    "fee_percent" int,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);