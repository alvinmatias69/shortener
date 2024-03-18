CREATE TABLE mst_url_shortener (
       id BIGINT PRIMARY KEY,
       hash VARCHAR(10) UNIQUE NOT NULL,
       long_url VARCHAR(150) UNIQUE NOT NULL
);
