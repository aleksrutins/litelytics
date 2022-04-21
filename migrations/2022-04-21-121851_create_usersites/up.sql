-- Your SQL goes here
CREATE TABLE usersites (
  id	serial	primary key,
  user_id integer references users(id) not null,
  site_id integer references sites(id) not null
);