-- Your SQL goes here

CREATE TABLE IF NOT EXISTS sites (
    id serial primary key,
    domain text unique not null
);