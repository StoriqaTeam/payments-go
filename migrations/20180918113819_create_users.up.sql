CREATE TABLE users (
    id VARCHAR(),
    email VARCHAR NOT NULL CHECK (email = lower(email)),
    email_verified BOOLEAN NOT NULL DEFAULT 'f',
    phone VARCHAR,
    phone_verified BOOLEAN NOT NULL DEFAULT 'f',
    is_active BOOLEAN NOT NULL DEFAULT 't',
    first_name VARCHAR,
    last_name VARCHAR,
    middle_name VARCHAR,
    gender VARCHAR,
    birthdate DATE,
    last_login_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp
);

CREATE UNIQUE INDEX users_email_idx ON users (email);

SELECT diesel_manage_updated_at('users');