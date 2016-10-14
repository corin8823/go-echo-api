
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name TEXT NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user;
