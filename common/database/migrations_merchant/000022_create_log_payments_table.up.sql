CREATE TABLE "log_payments" (
    "id" bigserial,
    "payment_id" bigint,
    "payment_trx_id" varchar(36),
    "merchant_id" int,
    "currency" varchar(20),
    "is_complete" boolean,
    "start_block" int,
    "tx_id" varchar(255),
    "tx_hash" varchar(255),
    "address" varchar(255),
    "data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);