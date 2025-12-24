CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    due_date TIMESTAMP WITH TIME ZONE NOT NULL,
    effort INTEGER CHECK (effort >= 1 AND effort <= 5),
    difficulty TEXT CHECK (difficulty IN ('Easy', 'Medium', 'Hard')),
    requirements TEXT[],
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
