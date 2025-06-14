
-- +migrate Up
ALTER TABLE plans ADD UNIQUE `idx_plans_supporter_id_project_id` (`supporter_id`, `project_id`);

-- +migrate Down
ALTER TABLE plans DROP INDEX `idx_plans_supporter_id_project_id`;
