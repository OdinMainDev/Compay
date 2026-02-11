CREATE TABLE "log_user" (
    "id" bigserial,
    "user_id" bigint,
    "merchant_id" int,
    "action" varchar(60),
    "url" text,
    "request" text,
    "client_data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_log_user_user_id ON "log_user" ("user_id");
CREATE INDEX idx_log_user_merchant_id ON "log_user" ("merchant_id");
CREATE INDEX idx_log_user_created_at ON "log_user" ("created_at" DESC);