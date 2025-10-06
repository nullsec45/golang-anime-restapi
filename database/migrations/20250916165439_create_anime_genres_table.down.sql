DROP TABLE IF EXISTS "anime_genres";

BEGIN;

DROP TRIGGER IF EXISTS trg_genres_many_anime_genres_soft_delete ON public.genres;
DROP TRIGGER IF EXISTS trg_animes_many_anime_genres_soft_delete  ON public.animes;

DROP FUNCTION IF EXISTS public.hard_delete_links_on_soft_delete_anime_genres();

COMMIT;