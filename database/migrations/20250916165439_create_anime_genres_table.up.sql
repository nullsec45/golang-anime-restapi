CREATE TABLE anime_genres (
  id UUID NOT NULL,
  anime_id UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE,
  genre_id UUID NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
  created_at    TIMESTAMPTZ NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NULL DEFAULT now(),
  PRIMARY KEY (id,anime_id, genre_id)
);