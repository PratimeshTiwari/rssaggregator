-- name: CreateCourse :one 
INSERT INTO courses (id, created_at, updated_at,name,url,user_id)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING *;

-- name: GetCourses :many 
SELECT * FROM courses;

-- name: GetNextCoursesToFetch :many
SELECT * FROM courses
ORDER BY last_fetched_at ASC NULLS FIRST 
LIMIT $1;

-- name: MarkCourseAsFetched :one
UPDATE courses
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;  