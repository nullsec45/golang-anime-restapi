CREATE TABLE voice_casts (
  id           UUID DEFAULT gen_random_uuid() NOT NULL,
  anime_id     UUID NOT NULL REFERENCES animes(id) ON DELETE CASCADE ON UPDATE CASCADE,
  character_id UUID NOT NULL REFERENCES characters(id) ON DELETE CASCADE ON UPDATE CASCADE,
  person_id    UUID NOT NULL REFERENCES peoples(id) ON DELETE CASCADE ON UPDATE CASCADE,
  language     TEXT NOT NULL DEFAULT 'Japanese',
  role_note    TEXT,                       
  created_at   TIMESTAMPTZ  NULL DEFAULT NULL,
  updated_at   TIMESTAMPTZ  NULL DEFAULT NULL,
  
  PRIMARY KEY (anime_id, character_id, person_id)
  
);

CREATE OR REPLACE FUNCTION hard_delete_links_on_soft_delete_voice_casts()
RETURNS trigger
LANGUAGE plpgsql
AS $$
BEGIN

  IF NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at) THEN
    IF TG_TABLE_NAME = 'animes' THEN
      DELETE FROM voice_casts WHERE anime_id = NEW.id;
    ELSIF TG_TABLE_NAME = 'characters' THEN
      DELETE FROM voice_casts WHERE character_id = NEW.id;
    ELSIF TG_TABLE_NAME = 'peoples' THEN
      DELETE FROM voice_casts WHERE person_id = NEW.id;
    END IF;
  END IF;
  RETURN NEW;
END;
$$;

DROP TRIGGER IF EXISTS trg_animes_many_voice_casts_soft_delete ON animes;
CREATE TRIGGER trg_animes_many_voice_casts_soft_delete
AFTER UPDATE OF deleted_at ON animes
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_voice_casts();

DROP TRIGGER IF EXISTS trg_animes_many_voice_casts_soft_delete ON characters;
CREATE TRIGGER trg_animes_many_voice_casts_soft_delete
AFTER UPDATE OF deleted_at ON characters
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_voice_casts();

DROP TRIGGER IF EXISTS trg_animes_many_voice_casts_soft_delete ON peoples;
CREATE TRIGGER trg_animes_many_voice_casts_soft_delete
AFTER UPDATE OF deleted_at ON peoples
FOR EACH ROW
WHEN (NEW.deleted_at IS NOT NULL AND (OLD.deleted_at IS DISTINCT FROM NEW.deleted_at))
EXECUTE FUNCTION hard_delete_links_on_soft_delete_voice_casts();