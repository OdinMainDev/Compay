CREATE TABLE "blockchain_states" (
    "id" bigserial,
    "currency" varchar(20),
    "current_block" bigint,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);