CREATE TABLE IF NOT EXISTS "public"."media" (
    id UUID DEFAULT gen_random_uuid() NOT NULL,
    path text,
    created_at TIMESTAMPTZ NULL DEFAULT NULL,
    updated_at TIMESTAMPTZ NULL DEFAULT NULL,
    CONSTRAINT "media_pk" PRIMARY KEY ("id")
);