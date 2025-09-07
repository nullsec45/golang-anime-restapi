CREATE TABLE "public"."users" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL
)