ALTER TABLE peoples
  DROP CONSTRAINT IF EXISTS peoples_person_image_fkey;

ALTER TABLE peoples
  DROP COLUMN IF EXISTS peoples_person_image;