CREATE TABLE "notifications" (
    "id" bigserial,
    "merchant_id" int,
    "is_read" boolean,
    "action" varchar(255),
    "text" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);