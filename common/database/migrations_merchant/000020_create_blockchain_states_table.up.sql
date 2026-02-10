CREATE TABLE "blockchain_states" (
    "id" bigserial,
    "currency" varchar(20),
    "current_block" bigint,
    "created_at" int,
    "updated_at" int,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX idx_blockchain_states_currency ON "blockchain_states" ("currency");