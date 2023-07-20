CREATE TABLE IF NOT EXISTS "habit" (
    "id" serial PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "start_date" date NOT NULL,
    "end_date" date,
    "streak_count" integer,
    "completed" boolean,
    "comments" text,
    "category" varchar(255)
);