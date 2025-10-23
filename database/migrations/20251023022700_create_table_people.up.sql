DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
    CREATE TYPE gender_enum AS ENUM ('Male','Female');
  END IF;
END$$;

CREATE TABLE IF NOT EXISTS peoples (
  id          UUID PRIMARY KEY DEFAULT gen_random_uuid() NOT NULL,
  slug        VARCHAR(255) NOT NULL UNIQUE,
  name_native VARCHAR(255),
  name        VARCHAR(255) NOT NULL UNIQUE,         
  birthday    DATE NOT NULL,
  gender      gender_enum NOT NULL,
  country     VARCHAR(150) NOT NULL,
  site_url    VARCHAR(200),
  biography   TEXT NOT NULL,
  created_at  TIMESTAMPTZ  NULL DEFAULT NULL,
  updated_at  TIMESTAMPTZ  NULL DEFAULT NULL,
  deleted_at  TIMESTAMPTZ  NULL DEFAULT NULL
);
