ALTER TABLE habit
    ADD COLUMN user_id INTEGER;

ALTER TABLE habit
    ADD CONSTRAINT fk_user_id
        FOREIGN KEY (user_id) REFERENCES "user"(id);