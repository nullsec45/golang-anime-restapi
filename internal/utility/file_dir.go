package utility

import (
	"path"
	"path/filepath"
	"strings"
	"github.com/nullsec45/golang-anime-restapi/domain"
)

func SafeJoin(baseDir,rel string) (absolute string, err error) {
	cleanRel := filepath.Clean(rel)
	abs := filepath.Join(baseDir, cleanRel)

	baseAbs, err := filepath.Abs(baseDir)

	if err != nil {
		return "", err
	}
	absTarget, err := filepath.Abs(abs)
	if err != nil {
		return "", err
	}

	sep := string(filepath.Separator)
	if !strings.HasPrefix(absTarget+sep, baseAbs+sep) && absTarget != baseAbs {
		return "", domain.AnimeMediaOutsideDir
	}

	return absTarget, nil
}

func PublicURL(baseURL, rel string) string {
	trimmed := strings.TrimRight(baseURL, "/")
	return path.Join(trimmed+"/", rel)
}
