CREATE TABLE users (
  id	serial	primary key,
  email	text	unique,
  password	text
);

CREATE TABLE sites (
  id	serial	primary key,
  domain text	unique,
  name	text
);

CREATE TABLE usersites (
  id	serial	primary key,
  user_id integer references users(id),
  site_id integer references sites(id)
);

CREATE TABLE visits (
  id	serial	primary key,
  site	integer	references sites(id),
  path	text,
  referer	text,
  timestamp	text,
  ip	text,
  useragent text
);