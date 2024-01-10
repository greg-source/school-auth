-- name: GetProfiles :many
select id, username, first_name, last_name, city, school from user
      RIGHT JOIN user_profile up on user.id = up.user_id
      RIGHT JOIN user_data ud on user.id = ud.user_id;
-- name: GetProfile :one
select id, username, first_name, last_name, city, school from user
      RIGHT JOIN user_profile up on user.id = up.user_id
      RIGHT JOIN user_data ud on user.id = ud.user_id
where username = ?;
-- name: CountApiKey :one
SELECT COUNT(*) AS count FROM auth WHERE api_key = ?;