-- Your SQL goes here

CREATE TABLE sites (
    id serial primary key,
    domain text unique not null
);