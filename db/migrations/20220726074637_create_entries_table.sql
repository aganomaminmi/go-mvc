-- migrate:up
create table users (
  id integer NOT NULL,
  first_name varchar(255),
  last_name varchar(255),
  email varchar(255) NOT NULL,
  age int(3),
  sex varchar(25),
  PRIMARY KEY (id)
)


-- migrate:down

