CREATE TABLE IF NOT EXISTS users {
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL
}

CREATE TABLE IF NOT EXISTS sessions {
    token UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    expire TIMESTAMP NOT NULL
}