DROP TABLE IF EXISTS anime_studios;

BEGIN;

-- 1) Hapus trigger
DROP TRIGGER IF EXISTS trg_studios_many_anime_studios_soft_delete ON public.studios;
DROP TRIGGER IF EXISTS trg_animes_many_anime_studios_soft_delete  ON public.animes;

-- 2) Hapus function (pastikan tidak ada trigger lain yang masih refer ke fungsi ini)
DROP FUNCTION IF EXISTS public.hard_delete_links_on_soft_delete_anime_studios();

COMMIT;