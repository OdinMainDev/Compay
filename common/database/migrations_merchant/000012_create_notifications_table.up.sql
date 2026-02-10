CREATE TABLE "notifications" (
    "id" bigserial,
    "merchant_id" int,
    "is_read" boolean,
    "action" varchar(20),
    "text" text,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_notifications_merchant_id ON "notifications" ("merchant_id");
CREATE INDEX idx_notifications_merchant_id_is_read ON "notifications" ("merchant_id", "is_read");