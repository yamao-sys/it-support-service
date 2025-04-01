
-- +migrate Up
ALTER TABLE projects modify start_date DATE NOT NULL;
ALTER TABLE projects modify end_date DATE NOT NULL;

-- +migrate Down
ALTER TABLE projects modify start_date DATETIME NOT NULL;
ALTER TABLE projects modify end_date DATETIME NOT NULL;
