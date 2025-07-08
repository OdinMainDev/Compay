CREATE TABLE "users" (
    "id" bigserial,
    "token" varchar(128),
    "pin_code" varchar(44),
    "is_banned" boolean,
    "secret_token" varchar(128),
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_users_token ON "users" ("token");
