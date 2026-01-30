CREATE TABLE "merchants" (
    "id" bigserial,
    "name" varchar(50),
    "token" varchar(36),
    "secret_key" varchar(255),
    "creator_id" bigint,
    "owner_id" bigint,
    "is_enabled" boolean,
    "settings" jsonb,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

