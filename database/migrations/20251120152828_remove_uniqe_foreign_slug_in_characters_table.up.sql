ALTER TABLE characters
  DROP CONSTRAINT IF EXISTS "characters_slug_key";

ALTER TABLE characters
  DROP CONSTRAINT IF EXISTS "characters_name_key";

DROP INDEX IF EXISTS "characters_slug_key";
DROP INDEX IF EXISTS "characters_name_key";