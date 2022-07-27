-- migrate:up

ALTER TABLE users
   modify first_name varchar(255) not null,
   modify last_name varchar(255) not null;


-- migrate:down

