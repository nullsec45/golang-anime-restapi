CREATE TABLE anime_tags (
  id UUID NOT NULL,
  anime_id UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE,
  tag_id UUID NOT NULL REFERENCES tags(id) ON DELETE CASCADE ON UPDATE CASCADE,
  created_at    TIMESTAMPTZ NULL DEFAULT NULL,
  updated_at    TIMESTAMPTZ NULL DEFAULT NULL,
  PRIMARY KEY (id,anime_id, tag_id)
);

CREATE OR REPLACE FUNCTION hard_delete_links_on_soft_delete_anime_tags()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN

  IF NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at) THEN
    IF TG_TABLE_NAME = 'tags' THEN
      DELETE FROM anime_tags WHERE tag_id = NEW.id;
    ELSIF TG_TABLE_NAME = 'animes' THEN
      DELETE FROM anime_tags WHERE tag_id = NEW.id;
    END IF;
  END IF;
  RETURN NEW;
END;
$$;

DROP TRIGGER IF EXISTS trg_tags_soft_delete ON tags;
CREATE TRIGGER trg_tags_many_anime_tags_soft_delete
AFTER UPDATE OF deleted_at ON tags
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_anime_tags();

DROP TRIGGER IF EXISTS trg_animes_soft_delete ON animes;
CREATE TRIGGER trg_animes_many_anime_tags_soft_delete
AFTER UPDATE OF deleted_at ON animes
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_anime_tags();