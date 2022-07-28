-- migrate:up
ALTER TABLE users
  ADD created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  ADD updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;


-- migrate:down
