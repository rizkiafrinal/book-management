-- +migrate Up
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
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR,
    modified_at TIMESTAMP,
    modified_by VARCHAR
);

-- +migrate Down
DROP TABLE IF EXISTS books;
