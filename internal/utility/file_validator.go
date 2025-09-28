package utility

import (
	"errors"
	"path/filepath"
	"strings"
	"mime/multipart"
)

func normalizeExt(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "" {
		return ""
	}

	if !strings.HasPrefix(s, ".") {
		s = "." + s
	}
	
	return s
}


func ValidateMediaFile(fh *multipart.FileHeader, allowedExts []string, maxSizeBytes int64) error {
	if fh == nil || fh.Filename == "" {
		return errors.New("'Media' files are required to be uploaded.")
	}
	ext := strings.ToLower(filepath.Ext(fh.Filename))

	ok := false
	for _, a := range allowedExts {
		if ext == normalizeExt(a) {
			ok = true
			break
		}
	}

	if !ok {
		return errors.New("Extensions are not allowed.")
	}
	
	if maxSizeBytes > 0 && fh.Size > maxSizeBytes {
		return errors.New("File size exceeds the limit.")
	}

	return nil
}