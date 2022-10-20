-- Add up migration script here
CREATE TABLE visits (
  id serial primary key,
  site integer references sites(id) not null,
  path text not null,
  referer text not null,
  timestamp text not null,
  ip text
);