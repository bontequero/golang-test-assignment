# Postgres

CREATE DATABASE test_db;
CREATE USER test_user WITH PASSWORD 'test_password';
GRANT ALL ON DATABASE test_db TO test_user;

CREATE TABLE groups (
  id smallint PRIMARY KEY,
  name varchar (20) NOT NULL
);

CREATE TABLE users (
  id serial PRIMARY KEY,
  group_id int REFERENCES groups(id),
  login varchar (64) NOT NULL,
  password varchar (64) NOT NULL
);

CREATE TABLE notes (
  id serial PRIMARY KEY,
  user_id int REFERENCES users(id) NOT NULL,
  title varchar (200) NOT NULL,
  content varchar
);

INSERT INTO groups (id, name)
VALUES (1, 'admin'), (2, 'user');

INSERT INTO users (group_id, login, password)
VALUES (1, 'Samanta', 'qw'), (2, 'Antony', 'er');
