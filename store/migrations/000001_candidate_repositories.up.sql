CREATE TABLE candidate_repositories (
    did VARCHAR(64) PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    is_artist BOOL NOT NULL,
    comment TEXT NOT NULL
);