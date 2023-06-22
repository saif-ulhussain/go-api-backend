CREATE TABLE habit (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
start_date DATE NOT NULL,
end_date DATE,
streak_count INTEGER,
completed BOOLEAN,
comments TEXT,
category VARCHAR(255)
);