CREATE TABLE pregnancy_access(
  id SERIAL PRIMARY KEY,
  pregnancy_id INTEGER NOT NULL REFERENCES pregnancies(id) ON DELETE CASCADE,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  role TEXT NOT NULL DEFAULT 'guesser',
  invited_by INTEGER REFERENCES users(id),
  access_token TEXT, 
  created_at TIMESTAMP DEFAULT NOW(),
  UNIQUE(pregnancy_id,user_id)
);
