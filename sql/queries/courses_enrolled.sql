-- name: CreateCoursesEnrolled :one 
INSERT INTO courses_enrolled (id, created_at, updated_at,user_id,course_id)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: GetCoursesEnrolled :many 
SELECT * FROM courses_enrolled WHERE user_id = $1;

