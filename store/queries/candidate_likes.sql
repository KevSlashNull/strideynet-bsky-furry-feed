-- name: CreateCandidateLike :exec
INSERT INTO
    candidate_likes (uri, actor_did, subject_uri, created_at,
                     indexed_at)
VALUES
    ($1, $2, $3, $4, $5);