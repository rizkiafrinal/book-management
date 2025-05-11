-- +migrate Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR,
    image_url VARCHAR,
    release_year INT CHECK (release_year BETWEEN 1980 AND 2024),
    price INT,
    total_page INT,
    thickness VARCHAR,
    category_id INT REFERENCES categories(id),
    created_at TIMESTAMP DEFAULT now(),
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);
-- +migrate Down
DROP TABLE books;
DROP TABLE categories;
DROP TABLE users;