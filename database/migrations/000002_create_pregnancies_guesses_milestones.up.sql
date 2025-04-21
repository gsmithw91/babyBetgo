--Pregnancies Table--
CREATE TABLE pregnancies(
  id SERIAL Primary KEY,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  due_date DATE NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()


);


--Guesses Table-- 
CREATE TABLE guesses (
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  pregnancy_id INTEGER NOT NULL REFERENCES pregnancies(id), 
  gender_guess VARCHAR(10) NOT NULL CHECK (gender_guess IN ('male','female','other')),
  created_at TIMESTAMP DEFAULT NOW()

);
--Milestones Table--

CREATE TABLE milestones(
  id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL REFERENCES pregnancies(id) ON DELETE CASCADE,
  week INTEGER NOT NULL CHECK (week >=1 AND week <= 42),
  title TEXT NOT NULL,
  description TEXT, 
  image_url TEXT,
  created_at TIMESTAMP DEFAULT NOW()
);




