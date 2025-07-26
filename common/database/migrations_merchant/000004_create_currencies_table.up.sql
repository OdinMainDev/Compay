CREATE TABLE "currencies" (
    "id" bigserial,
    "code_alpha" varchar(20),
    "title" varchar(255),
    "image" varchar(255),
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_currencies_code_alpha ON "currencies" ("code_alpha");
