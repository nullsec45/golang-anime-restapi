ALTER TABLE characters
  DROP CONSTRAINT IF EXISTS characters_character_image_fkey;

ALTER TABLE characters
  DROP COLUMN IF EXISTS character_image;