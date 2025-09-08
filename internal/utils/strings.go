package utils

import "strings"

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func Truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func Pad(s string, n int) string {
	for len(s) < n {
		s += " "
	}
	return s
}
