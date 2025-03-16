-- +goose Up
-- +goose StatementBegin
-- Миграция для таблицы users
CREATE TABLE users (
  id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  role varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

-- Миграция для таблицы categories
CREATE TABLE categories (
  id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
  name varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

-- Миграция для таблицы food
CREATE TABLE food (
  id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
  name varchar(255) NOT NULL,
  description varchar(255),
  cost integer NOT NULL,
  category_id integer NOT NULL,
  picture varchar(255),
  PRIMARY KEY (id),
  CONSTRAINT fk_food_category FOREIGN KEY (category_id)
    REFERENCES categories (id)
);

-- Миграция для таблицы tables
CREATE TABLE tables (
  id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
  seat integer NOT NULL,
  number_seats integer NOT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id)
);

-- Миграция для таблицы bookings
CREATE TABLE bookings (
  id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
  id_table integer NOT NULL,
  date date NOT NULL,
  time time NOT NULL,
  full_name varchar(255) NOT NULL,
  phone varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  count_seats integer NOT NULL,
  number_of_people integer NOT NULL,
  status varchar(255) NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_booking_table FOREIGN KEY (id_table)
    REFERENCES tables (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS bookings CASCADE;
DROP TABLE IF EXISTS tables CASCADE;
DROP TABLE IF EXISTS food CASCADE;
DROP TABLE IF EXISTS categories CASCADE;
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
