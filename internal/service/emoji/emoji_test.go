package emoji

import "testing"

func TestGetEmoji(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Folder 😀", "😀"},
		{"File.md", ""},
		{"Image.png 📷", "📷"},
		{"NoEmojiHere", ""},
	}

	for _, tt := range tests {
		result := GetEmoji(tt.name)
		if result != tt.expected {
			t.Errorf("GetEmoji(%q) = %q; want %q", tt.name, result, tt.expected)
		}
	}
}

func TestAddEmoji(t *testing.T) {
	tests := []struct {
		name     string
		emoji    string
		expected string
	}{
		{"Folder", "😀", "Folder 😀"},
		{"Folder 😀", "😀", "Folder 😀"},
		{"File.md", "📄", "File.md 📄"},
		{"File.md 📄", "📄", "File.md 📄"},
	}

	for _, tt := range tests {
		result := AddEmoji(tt.name, tt.emoji)
		if result != tt.expected {
			t.Errorf("AddEmoji(%q, %q) = %q; want %q", tt.name, tt.emoji, result, tt.expected)
		}
	}
}
