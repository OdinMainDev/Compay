CREATE TABLE "users" (
    "id" bigserial,
    "merchant_id" int,
    "tg_id" bigint,
    "first_name" varchar(255),
    "last_name" varchar(255),
    "username" varchar(255),
    "is_premium" boolean DEFAULT false,
    "is_active" boolean DEFAULT false,
    "is_approved" boolean DEFAULT false,
    "secret_token" varchar(255) NOT NULL,
    "totp_token" varchar(255) NOT NULL,
    "is_twofa_enabled" boolean DEFAULT false,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_users_tg_id ON "users" ("tg_id");
CREATE INDEX idx_users_username ON "users" ("username");
CREATE INDEX idx_users_is_approved ON "users" ("is_approved");
CREATE INDEX idx_users_is_twofa_enabled ON "users" ("is_twofa_enabled");
CREATE INDEX idx_users_is_premium ON "users" ("is_premium");
CREATE INDEX idx_users_is_active ON "users" ("is_active");