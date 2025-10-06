CREATE TABLE anime_genres (
  id UUID NOT NULL,
  anime_id UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE,
  genre_id UUID NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
  created_at    TIMESTAMPTZ NULL DEFAULT NULL,
  updated_at    TIMESTAMPTZ NULL DEFAULT NULL,
  PRIMARY KEY (id,anime_id, genre_id)
);

CREATE OR REPLACE FUNCTION hard_delete_links_on_soft_delete_anime_genres()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN

  IF NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at) THEN
    IF TG_TABLE_NAME = 'genres' THEN
      DELETE FROM anime_genres WHERE genre_id = NEW.id;
    ELSIF TG_TABLE_NAME = 'animes' THEN
      DELETE FROM anime_genres WHERE anime_id = NEW.id;
    END IF;
  END IF;
  RETURN NEW;
END;
$$;

DROP TRIGGER IF EXISTS trg_genres_many_anime_genres_soft_delete ON genres;
CREATE TRIGGER trg_genres_many_anime_genres_soft_delete
AFTER UPDATE OF deleted_at ON genres
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_anime_genres();

DROP TRIGGER IF EXISTS trg_animes_many_anime_genres_soft_delete ON animes;
CREATE TRIGGER trg_animes_many_anime_genres_soft_delete
AFTER UPDATE OF deleted_at ON animes
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_anime_genres();