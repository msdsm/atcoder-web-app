CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    atcoder_id VARCHAR(255) NOT NULL
);

CREATE TABLE rivals (
    id UUID PRIMARY KEY,
    rival_atcoder_id VARCHAR(255) NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    CONSTRAINT unique_user_rival UNIQUE (user_id, rival_atcoder_id)
);