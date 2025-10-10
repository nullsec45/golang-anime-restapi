package utility

import (
	"strings"
	"unicode/utf8"
)


func MaskString(s string, keepLeft, keepRight int) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "<empty>"
	}
	runes := []rune(s)
	n := len(runes)
	if keepLeft+keepRight >= n {
		return strings.Repeat("*", n)
	}
	left := string(runes[:keepLeft])
	right := string(runes[n-keepRight:])
	return left + strings.Repeat("*", n-keepLeft-keepRight) + right
}

func MaskEmail(email string) string {
	email = strings.TrimSpace(email)
	if email == "" {
		return "<empty>"
	}
	at := strings.Index(email, "@")
	if at <= 0 || at == len(email)-1 {
		return MaskString(email, 1, 1)
	}
	local := email[:at]
	domain := email[at+1:]

	localMasked := local
	if utf8.RuneCountInString(local) >= 3 {
		localMasked = MaskString(local, 1, 1)
	} else {
		localMasked = strings.Repeat("*", len(local))
	}
	
	return localMasked + "@" + domain
}

func MaskPhone(phone string) string {
	phone = strings.TrimSpace(phone)
	if phone == "" {
		return "<empty>"
	}
	prefix := ""
	if strings.HasPrefix(phone, "+") {
		prefix = "+"
		phone = strings.TrimPrefix(phone, "+")
	}
	return prefix + MaskString(phone, 2, 2)
}
