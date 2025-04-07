
-- +migrate Up
ALTER TABLE projects modify is_active BOOLEAN NOT NULL;

-- +migrate Down
ALTER TABLE projects modify is_active BOOLEAN NOT NULL DEFAULT TRUE;
