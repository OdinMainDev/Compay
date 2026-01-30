CREATE TABLE "log_auth" (
    "id" bigserial,
    "user_id" bigint,
    "merchant_id" int,
    "action" varchar(20),
    "url" text,
    "ip" varchar(20),
    "fingerprint" varchar(255),
    "user_agent" varchar(255),
    "client_data" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);