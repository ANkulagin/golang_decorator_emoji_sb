package emoji

import (
	"regexp"
)

// Регулярное выражение для обнаружения эмодзи.
var emojiRegex = regexp.MustCompile(`[\x{1F300}-\x{1F5FF}]|[\x{1F600}-\x{1F64F}]|[\x{1F680}-\x{1F6FF}]|[\x{1F700}-\x{1F77F}]|[\x{1F900}-\x{1F9FF}]|[\x{2600}-\x{26FF}]|[\x{2700}-\x{27BF}]|[\x{1FA70}-\x{1FAFF}]|[\x{1F1E6}-\x{1F1FF}]`)

// GetEmoji извлекает первый найденный эмодзи из строки.
func GetEmoji(name string) string {
	match := emojiRegex.FindString(name)
	return match
}

// AddEmoji добавляет эмодзи к имени, если его там нет.
func AddEmoji(name, emoji string) string {
	if !ContainsEmoji(name) {
		return name + " " + emoji
	}
	return name
}

// ContainsEmoji проверяет, есть ли в строке эмодзи.
func ContainsEmoji(name string) bool {
	return emojiRegex.MatchString(name)
}
