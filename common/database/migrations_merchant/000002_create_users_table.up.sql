CREATE TABLE "users" (
    "id" bigserial,
    "merchant_id" int,
    "username" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "secret_token" varchar(255) NOT NULL,
    "totp_token" varchar(255) NOT NULL,
    "is_twofa_enabled" boolean DEFAULT false,
    "is_banned" boolean DEFAULT false,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_users_merchant_id ON "users" ("merchant_id");
CREATE INDEX idx_users_username ON "users" ("username");