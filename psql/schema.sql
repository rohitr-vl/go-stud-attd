CREATE TABLE users (
  ID   BIGSERIAL PRIMARY KEY,
  name text      NOT NULL,
  email text     NOT NULL UNIQUE,
  bio  text,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE tasks (
    ID BIGSERIAL PRIMARY KEY,
    title text NOT NULL,
    description text NOT NULL,
    created_by_user_id BIGINT NOT NULL,
    assigned_to_user_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);