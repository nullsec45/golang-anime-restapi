DROP TABLE IF EXISTS voice_casts;

BEGIN;

DROP TRIGGER IF EXISTS trg_animes_many_voice_casts_soft_delete  ON public.animes;
DROP TRIGGER IF EXISTS trg_genres_many_voice_casts_soft_delete ON public.characters;
DROP TRIGGER IF EXISTS trg_genres_many_voice_casts_soft_delete ON public.peoples;

DROP FUNCTION IF EXISTS public.hard_delete_links_on_soft_delete_voice_casts();

COMMIT;