CREATE TABLE IF NOT EXISTS users(
  user_id serial PRIMARY KEY,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  username varchar(50) UNIQUE NOT NULL,
  password varchar(60) NOT NULL,
  email varchar(300) UNIQUE
);
