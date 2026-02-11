CREATE TABLE "log_auth" (
    "id" bigserial,
    "user_id" bigint,
    "merchant_id" int,
    "action" varchar(60),
    "url" text,
    "ip" varchar(20),
    "fingerprint" varchar(255),
    "user_agent" varchar(255),
    "client_data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_log_auth_user_id ON "log_auth" ("user_id");
CREATE INDEX idx_log_auth_merchant_id ON "log_auth" ("merchant_id");
CREATE INDEX idx_log_auth_created_at ON "log_auth" ("created_at" DESC);