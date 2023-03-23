CREATE TABLE IF NOT EXISTS incident(
       id serial PRIMARY KEY,
       user_id bigint unsigned,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP
    );