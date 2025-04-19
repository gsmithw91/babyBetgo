CREATE TABLE users (
  id                  SERIAL PRIMARY KEY,
  username            TEXT UNIQUE NOT NULL,
  password_hash       TEXT NOT NULL,
  email               TEXT UNIQUE NOT NULL,
  created_at          TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at          TIMESTAMP,
  is_active           BOOLEAN DEFAULT TRUE,
  last_login          TIMESTAMP,
  profile_picture_url TEXT,
  role                TEXT,
  display_name        TEXT,
  bio                 TEXT,
  phone_number        TEXT,
  balance             BIGINT NOT NULL DEFAULT 0
);

