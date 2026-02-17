SET search_path TO app;

CREATE TABLE IF NOT EXISTS schema_migrations (
    id SERIAL PRIMARY KEY,
	filename VARCHAR(255) UNIQUE NOT NULL,
	applied_at TIMESTAMP NOT NULL DEFAULT NOW()
);