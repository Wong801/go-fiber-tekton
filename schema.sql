CREATE TABLE "public"."mst_banks" (
    "bank_id" uuid NOT NULL DEFAULT uuid_generate_v4(),
    "bank_name" varchar NOT NULL,
    "bank_code" varchar NOT NULL,
    PRIMARY KEY ("bank_id")
);