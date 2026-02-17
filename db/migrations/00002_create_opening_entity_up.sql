CREATE TABLE IF NOT EXISTS opening (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    role TEXT NOT NULL,
	company TEXT NOT NULL,
    location TEXT,
    remote BOOLEAN NOT NULL DEFAULT FALSE,
    link TEXT,
    salary BIGINT
);