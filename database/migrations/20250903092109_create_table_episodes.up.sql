CREATE TABLE episodes (
  id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  anime_id      UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE,
  number        INT NOT NULL,                                
  season_number INT,                                          
  title         TEXT,
  synopsis      TEXT,
  air_date      DATE,
  duration_minutes INT,
  is_special    BOOLEAN NOT NULL DEFAULT false,
  created_at    TIMESTAMPTZ NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NULL DEFAULT now(),
  UNIQUE (anime_id, number)                                   
);