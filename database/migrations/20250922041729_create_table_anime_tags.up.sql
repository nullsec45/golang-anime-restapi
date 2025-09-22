CREATE TABLE anime_tags (
  id UUID NOT NULL,
  anime_id UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE,
  tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
  created_at    TIMESTAMPTZ NULL DEFAULT now(),
  updated_at    TIMESTAMPTZ NULL DEFAULT now(),
  PRIMARY KEY (id,anime_id, tag_id)
);