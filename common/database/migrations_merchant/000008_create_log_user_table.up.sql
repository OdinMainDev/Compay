CREATE TABLE "log_user" (
    "id" bigserial,
    "user_id" bigint,
    "merchant_id" int,
    "action" varchar(20),
    "url" text,
    "request" text,
    "client_data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);