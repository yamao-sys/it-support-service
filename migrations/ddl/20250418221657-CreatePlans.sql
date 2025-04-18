
-- +migrate Up
CREATE TABLE IF NOT EXISTS plans(
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	supporter_id INT NOT NULL,
	project_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	start_date DATE NOT NULL,
	end_date DATE NOT NULL,
	unit_price INT,
	status INT NOT NULL DEFAULT 0,
	agreed_at DATETIME,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (supporter_id) REFERENCES supporters(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS plans;
