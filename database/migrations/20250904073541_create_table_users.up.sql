CREATE TABLE "public"."users" (
    "id" character varying(36) DEFAULT 'gen_random_uuid()' NOT NULL,
    "email" character varying(255) NOT NULL,
    "password" character varying(255) NOT NULL
)