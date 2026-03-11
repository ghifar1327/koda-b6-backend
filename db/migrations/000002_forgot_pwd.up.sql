CREATE TABLE if NOT EXISTS forgot_password (
    id SERIAL PRIMARY KEY,
    email VARCHAR(10)  NOT NULL,
    code int NOT NULL
)
