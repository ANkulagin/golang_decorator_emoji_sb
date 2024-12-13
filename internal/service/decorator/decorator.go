package decorator

import (
	"github.com/ANkulagin/golang_decorator_emoji_sb/internal/service/emoji"
	"path/filepath"
)

type Decorator struct {
	Path  string
	Emoji string
}

func NewDecorator(path string) *Decorator {
	return &Decorator{
		Path: path,
	}
}

func (d *Decorator) DecorateDirectories() error {
	if d.Emoji == "" {
		d.Emoji = emoji.GetEmoji(filepath.Base(d.Path))
	}

	return d.walkDirectories() //d.Path, d.Emoji
}

func (d *Decorator) walkDirectories() error {
	return nil
}
