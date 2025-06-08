
-- +migrate Up
CREATE TABLE IF NOT EXISTS plan_steps(
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	plan_id INT NOT NULL,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	duration INT NOT NULL,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE IF EXISTS plan_steps;
