-- Your SQL goes here
CREATE TABLE IF NOT EXISTS users (
  id serial primary key,
  email text unique not null,
  password text not null
);