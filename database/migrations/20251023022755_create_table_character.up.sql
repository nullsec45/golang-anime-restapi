CREATE TABLE characters (
  id         UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  slug       VARCHAR(255) NOT NULL UNIQUE,
  name       VARCHAR(255) NOT NULL UNIQUE,
  name_native VARCHAR(255),
  description TEXT,
  created_at TIMESTAMPTZ  NULL DEFAULT NULL,
  updated_at TIMESTAMPTZ  NULL DEFAULT NULL,
  deleted_at TIMESTAMPTZ  NULL DEFAULT NULL
);