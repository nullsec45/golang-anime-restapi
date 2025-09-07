CREATE TABLE genres (
  id    UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  slug  TEXT NOT NULL UNIQUE,
  name  TEXT NOT NULL UNIQUE
);
