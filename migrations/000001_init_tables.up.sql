CREATE TABLE IF NOT EXISTS prices (
                        id SERIAL PRIMARY KEY,
                        coin VARCHAR(10) NOT NULL,
                        price DECIMAL(18, 8) NOT NULL,
                        timestamp BIGINT NOT NULL
);