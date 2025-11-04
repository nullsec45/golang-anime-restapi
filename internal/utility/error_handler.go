package utility

import (
	"fmt"
	"github.com/nullsec45/golang-anime-restapi/domain"
)

func NewNotFound(resource string) error {
	return fmt.Errorf("%s %w", resource, domain.ErrNotFound)
}

func NewAlreadyExist(resource string) error {
	return fmt.Errorf("%s %w", resource, domain.ErrAlreadyExist)
}

func NewAuthFailed(reason string) error {
	if reason == "" {
		return domain.ErrAuthFailed
	}

	return fmt.Errorf("%w: %s", domain.ErrAuthFailed, reason)
}
