CREATE TABLE IF NOT EXISTS "patient"(
    "id" SERIAL PRIMARY KEY,
    "hospital_id" INTEGER REFERENCES hospital(id),
    "full_name" VARCHAR(50) NOT NULL,
    "patient_info" VARCHAR(100) NOT NULL,
    "phone_number" VARCHAR(50) NOT NULL
);