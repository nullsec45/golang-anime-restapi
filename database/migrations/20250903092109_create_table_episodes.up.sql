CREATE TABLE episodes (
  id            BIGSERIAL PRIMARY KEY,
  anime_id      BIGINT NOT NULL REFERENCES anime(id) ON DELETE CASCADE,
  number        INT NOT NULL,                                  -- nomor episode (1..n)
  season_number INT,                                           -- jika ingin grouping per season
  title         TEXT,
  synopsis      TEXT,
  air_date      DATE,
  duration_minutes INT,
  is_special    BOOLEAN NOT NULL DEFAULT false,
  created_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE (anime_id, number)                                    -- tak boleh duplikat nomor
);