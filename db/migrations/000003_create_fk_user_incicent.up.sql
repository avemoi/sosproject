
ALTER TABLE incident
    ADD CONSTRAINT incident_user
        FOREIGN KEY (user_id)
            REFERENCES users (id);