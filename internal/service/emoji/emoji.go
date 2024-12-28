package emoji

import (
	"regexp"
	"strings"
)

// GetEmoji извлекает эмодзи из имени папки или файла, если оно есть.
func GetEmoji(name string) string {
	re := regexp.MustCompile(`\p{So}$`)
	match := re.FindString(name)
	return match
}

// AddEmoji добавляет эмодзи к имени, если его там нет.
func AddEmoji(name, emoji string) string {
	if !strings.HasSuffix(name, emoji) {
		return name + " " + emoji
	}
	return name
}

func ContainsEmoji(name string) bool {
	re := regexp.MustCompile(`\p{So}`)
	return re.MatchString(name)
}
