-- name: GetBioForAuthor :one
SELECT * FROM users WHERE id = ?;

-- name: CreateUser :execresult
INSERT INTO users (id, name) VALUES (?, ?);

-- name: GetAllGoals :many
SELECT * FROM goals WHERE user_id = ?;

-- name: GetGoalsWithName :many
SELECT u.name as UserName, g.name as GoalName, g.active FROM goals g, users u WHERE g.user_id = u.id AND u.name LIKE ?; 

-- name: GetDays :many
SELECT day FROM relapsed_days WHERE goal_id = ?;

-- name: MarkRelapse :exec
INSERT INTO relapsed_days (day, goal_id) VALUES (?, ?);

-- name: CreateGoal :execresult
INSERT INTO goals (user_id, name, content, active) VALUES (?, ?, ?, ?);

-- name: GetAllUsers :many
SELECT * FROM users;    

-- name: FetchFirstGoalDate :one
SELECT g.created_at FROM goals g WHERE g.user_id = ? ORDER BY g.created_at ASC LIMIT 1;

-- name: GetProgress :one
SELECT g.id, g.created_at, COUNT(r.day) fROM relapsed_days r, goals g WHERE g.id = r.goal_id AND g.id = ? GROUP BY g.id, g.created_at;
