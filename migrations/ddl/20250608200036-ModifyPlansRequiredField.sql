-- +migrate Up
ALTER TABLE plans modify start_date DATE;
ALTER TABLE plans modify end_date DATE;
ALTER TABLE plans modify unit_price INT NOT NULL DEFAULT 0;

-- +migrate Down
ALTER TABLE plans modify start_date DATE NOT NULL;
ALTER TABLE plans modify end_date DATE NOT NULL;
ALTER TABLE plans modify unit_price INT;
