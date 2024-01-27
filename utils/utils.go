package utils

import "strings"

func CreateSlug(str string) string {
	str = strings.ToLower(str)
	replacer := strings.NewReplacer(" ", "-", ";", "-", "/", "-", "\\", "-", "]", "-", "[", "-", "{", "-", "}", "-", "_", "-")
	str = replacer.Replace(str)
	return strings.Trim(str, "-")
}
