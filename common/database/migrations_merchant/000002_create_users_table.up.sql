CREATE TABLE "users" (
    "id" bigserial,
    "merchant_id" int,
    "secret_token" varchar(255) NOT NULL,
    "totp_token" varchar(255) NOT NULL,
    "is_twofa_enabled" boolean DEFAULT false,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_users_merchant_id ON "users" ("merchant_id");
CREATE INDEX idx_users_is_twofa_enabled ON "users" ("is_twofa_enabled");