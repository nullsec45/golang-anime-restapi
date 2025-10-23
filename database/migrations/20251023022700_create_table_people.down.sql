DROP TABLE IF EXISTS peoples;

DO $$
BEGIN
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_enum') THEN
    DROP TYPE gender_enum;
  END IF;
END$$;