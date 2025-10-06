DROP TABLE IF EXISTS "anime_tags";

BEGIN;

DROP TRIGGER IF EXISTS trg_tags_many_anime_tags_soft_delete ON public.tags;
DROP TRIGGER IF EXISTS trg_animes_many_anime_tags_soft_delete  ON public.animes;

DROP FUNCTION IF EXISTS public.hard_delete_links_on_soft_delete_anime_tags();

COMMIT;