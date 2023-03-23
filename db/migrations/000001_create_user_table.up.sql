CREATE TABLE IF NOT EXISTS users (
                                     id serial PRIMARY KEY,
                                     power_id int NOT NULL,
                                     latitude double precision NOT NULL,
                                     longitude double precision NOT NULL,
                                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
    );
