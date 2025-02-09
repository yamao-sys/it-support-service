
-- +migrate Up
ALTER TABLE supporters ADD birthday DATE AFTER password;

-- +migrate Down
ALTER TABLE supporters DROP COLUMN birthday;
