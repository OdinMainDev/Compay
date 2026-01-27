CREATE TABLE "merchants" (
    "id" bigserial,
    "name" varchar(255),
    "token" varchar(255),
    "creator_id" bigint,
    "owner_id" bigint,
    "is_enabled" boolean,
    "settings" jsonb,
    "created_at" int,
    "updated_at" int,
    "deleted_at" timestamptz,
    PRIMARY KEY ("id")
);

