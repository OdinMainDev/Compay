CREATE TABLE "currencies" (
    "id" bigserial,
    "code_alpha" varchar(20),
    "is_enabled" boolean,
    "title" varchar(255),
    "network" varchar(255),
    "image" varchar(255),
    "is_coin" boolean,
    "minor_units" smallint,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE INDEX idx_currencies_code_alpha ON "currencies" ("code_alpha");
