-- name: SingleChirp :one
 SELECT * FROM chirps
 WHERE ID = $1;