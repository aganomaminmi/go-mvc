-- migrate:up

ALTER TABLE users
  modify id int auto_increment;

-- migrate:down

