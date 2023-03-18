ALTER TABLE incident
    ADD CONSTRAINT unique_name
        UNIQUE (power_id);

