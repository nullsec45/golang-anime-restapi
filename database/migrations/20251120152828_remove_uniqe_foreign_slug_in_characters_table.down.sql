-- down.sql
ALTER TABLE characters
  ADD CONSTRAINT characters_slug_key UNIQUE (slug);

  ALTER TABLE characters
  ADD CONSTRAINT characters_slug_key UNIQUE (name);