CREATE TABLE IF NOT EXISTS anime_studios (
  id        UUID DEFAULT gen_random_uuid() NOT NULL,
  anime_id  UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE,
  studio_id UUID NOT NULL REFERENCES studios(id) ON DELETE CASCADE ON UPDATE CASCADE,
  role      TEXT NOT NULL CHECK (role IN ('Animation','Producer','Licensor','Publisher')),
  created_at TIMESTAMPTZ  NULL DEFAULT NULL,
  updated_at TIMESTAMPTZ  NULL DEFAULT NULL,
  PRIMARY KEY (id, anime_id, studio_id)
);

CREATE OR REPLACE FUNCTION hard_delete_links_on_soft_delete()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN
  -- Jalan hanya ketika deleted_at berubah dari NULL -> NOT NULL
  IF NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at) THEN
    IF TG_TABLE_NAME = 'studios' THEN
      DELETE FROM anime_studios WHERE studio_id = NEW.id;
    ELSIF TG_TABLE_NAME = 'animes' THEN
      DELETE FROM anime_studios WHERE anime_id = NEW.id;
    END IF;
  END IF;
  RETURN NEW;
END;
$$;

DROP TRIGGER IF EXISTS trg_studios_soft_delete ON studios;
CREATE TRIGGER trg_studios_soft_delete
AFTER UPDATE OF deleted_at ON studios
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete();

DROP TRIGGER IF EXISTS trg_animes_soft_delete ON animes;
CREATE TRIGGER trg_animes_soft_delete
AFTER UPDATE OF deleted_at ON animes
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete();