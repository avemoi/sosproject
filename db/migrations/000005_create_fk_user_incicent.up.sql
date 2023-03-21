ALTER TABLE incident
    ADD COLUMN user_id bigint unsigned;


ALTER TABLE incident
    ADD CONSTRAINT incident_user
        FOREIGN KEY (user_id)
            REFERENCES users (id);