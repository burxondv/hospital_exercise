CREATE TABLE IF NOT EXISTS "staff"(
    "id" SERIAL PRIMARY KEY,
    "hospital_id" INTEGER REFERENCES hospital(id),
    "full_name" VARCHAR(50) NOT NULL,
    "phone_number" VARCHAR NOT NULL
);