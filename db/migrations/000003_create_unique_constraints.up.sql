ALTER TABLE incident_type
    ADD CONSTRAINT unique_name
        UNIQUE (name);

