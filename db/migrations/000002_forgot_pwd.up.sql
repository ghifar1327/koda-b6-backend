CREATE TABLE if NOT EXISTS forgot_password (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    code int NOT NULL
);
