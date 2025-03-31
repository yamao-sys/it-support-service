
-- +migrate Up
CREATE TABLE IF NOT EXISTS projects(
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	company_id INT NOT NULL,
	title VARCHAR(255) NOT NULL,
	description TEXT NOT NULL,
	start_date DATETIME NOT NULL,
	end_date DATETIME NOT NULL,
	min_budget INT,
	max_budget INT,
	is_active BOOLEAN NOT NULL DEFAULT TRUE,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS projects;
