CREATE TABLE IF NOT EXISTS incident(
       id serial PRIMARY KEY,
       latitude double precision NOT NULL,
       longtitude double precision NOT NULL,
       incident_type_id BIGINT UNSIGNED not NULL,
       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
       FOREIGN KEY (incident_type_id) REFERENCES incident_type(id)
    );