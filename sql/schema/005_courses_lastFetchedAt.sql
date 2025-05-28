-- +goose Up
ALTER TABLE courses
ADD COLUMN last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE courses
DROP COLUMN last_fetched_at;
