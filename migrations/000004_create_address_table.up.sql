CREATE TABLE IF NOT EXISTS "address"(
    "id" SERIAL PRIMARY KEY,
    "hospital_id" INTEGER REFERENCES hospital(id),
    "region" VARCHAR(50) NOT NULL,
    "street" VARCHAR(50) NOT NULL
);