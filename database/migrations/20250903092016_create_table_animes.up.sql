-- CREATE EXTENSION IF NOT EXISTS pg_trgm;
-- CREATE EXTENSION IF NOT EXISTS unaccent;

CREATE TABLE animes (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  slug          TEXT NOT NULL UNIQUE,                          
  title_romaji  TEXT NOT NULL,
  title_native  TEXT,                                          
  title_english TEXT,
  synopsis      TEXT,
  type          TEXT NOT NULL CHECK (type IN ('TV','Movie','OVA','ONA','Special')),
  season        TEXT CHECK (season IN ('Winter','Spring','Summer','Fall')),
  season_year   SMALLINT CHECK (season_year BETWEEN 1917 AND 2100),
  status        TEXT NOT NULL CHECK (status IN ('Upcoming','Airing','Finished','Hiatus')),
  age_rating    TEXT CHECK (age_rating IN ('G','PG','PG-13','R','R+','Rx')),
  total_episodes INT,                                          
  average_duration_minutes INT,                                
  country      TEXT DEFAULT 'JP',
  premiered_at DATE,                                           
  ended_at     DATE,
  popularity   INT DEFAULT 0,                                  
  score_avg    NUMERIC(3,2),                                   
  alt_titles   JSONB DEFAULT '{}'::jsonb,                      
  external_ids JSONB DEFAULT '{}'::jsonb,                      
  created_at   TIMESTAMPTZ  NULL DEFAULT NULL,
  updated_at   TIMESTAMPTZ  NULL DEFAULT NULL,
  deleted_at   TIMESTAMPTZ  NULL 
);


-- CREATE INDEX idx_anime_title_trgm
--   ON anime USING gin ((unaccent(lower(coalesce(title_romaji,'') || ' ' || coalesce(title_english,'') || ' ' || coalesce(title_native,'')))) gin_trgm_ops);

-- CREATE INDEX idx_anime_external_ids_gin ON anime USING gin (external_ids);

