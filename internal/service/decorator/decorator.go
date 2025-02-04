package decorator

import (
	"fmt"
	"github.com/ANkulagin/golang_decorator_emoji_sb/internal/service/emoji"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Decorator struct {
	Path             string
	ConcurrencyLimit int
	skipPatterns     []string
}

func NewDecorator(path string, concurrencyLimit int, skipPatterns []string) *Decorator {
	return &Decorator{
		Path:             path,
		ConcurrencyLimit: concurrencyLimit,
		skipPatterns:     skipPatterns,
	}
}

func (d *Decorator) Decorate() error {
	emoji := emoji.GetEmoji(filepath.Base(d.Path))

	wg := &sync.WaitGroup{}
	sem := make(chan struct{}, d.ConcurrencyLimit)

	wg.Add(1)
	go d.decorateDirConcurrently(d.Path, emoji, wg, sem)
	wg.Wait()

	return nil
}

func (d *Decorator) decorateDirConcurrently(rootPath, emojiPath string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()

	sem <- struct{}{}
	defer func() {
		<-sem
	}()

	err := d.processDirectory(rootPath, emojiPath, wg, sem)
	if err != nil {
		log.WithError(err).Warnf("Ошибка при обработке директории: %s", rootPath)
	}
}

func (d *Decorator) processDirectory(rootPath, emojiPath string, wg *sync.WaitGroup, sem chan struct{}) error {
	dirBase := filepath.Base(rootPath)

	// Если имя директории начинается с одного из указанных префиксов, пропускаем её.
	if d.shouldSkipDirectory(dirBase) {
		log.Infof("Пропуск директории: %s (совпадает с шаблоном)", rootPath)
		return nil
	}

	dirEmoji := emoji.GetEmoji(dirBase)

	if dirEmoji == "" && emojiPath != "" {
		newName := emoji.AddEmoji(dirBase, emojiPath)
		if newName != dirBase {
			newPath := filepath.Join(filepath.Dir(rootPath), newName)
			if err := os.Rename(rootPath, newPath); err != nil {
				return fmt.Errorf("не удалось переименовать: %s -> %s %v", rootPath, newPath, err.Error())
			}
			rootPath = newPath
			dirEmoji = emojiPath
		}
	}

	entries, err := os.ReadDir(rootPath)
	if err != nil {
		return fmt.Errorf("не удалось прочитать директорию: %s %v", rootPath, err.Error())
	}

	for _, entry := range entries {
		oldName := entry.Name()
		oldPath := filepath.Join(rootPath, oldName)

		if entry.IsDir() {
			wg.Add(1)
			go d.decorateDirConcurrently(oldPath, dirEmoji, wg, sem)
		} else {
			err = addEmojiToFilename(oldName, dirEmoji, rootPath, oldPath)
			if err != nil {
				return fmt.Errorf("%v", err)
			}
		}
	}

	return nil
}

// shouldSkipDirectory возвращает true, если имя директории начинается с одного из указанных префиксов.
func (d *Decorator) shouldSkipDirectory(directoryName string) bool {
	for _, prefix := range d.skipPatterns {
		if strings.HasPrefix(directoryName, prefix) {
			return true
		}
	}
	return false
}

func addEmojiToFilename(oldName, inheritedEmoji, path, oldPath string) error {
	// Пропускаем файлы, не являющиеся .md
	if filepath.Ext(oldName) != ".md" {
		return nil
	}

	if emoji.ContainsEmoji(oldName) {
		return nil
	}

	// Разделяем имя файла и проверяем на наличие эмодзи
	fileParts := strings.SplitN(oldName, " ", 2)
	if len(fileParts) < 2 {
		// Пропускаем, если формат не соответствует ожиданиям
		return nil
	}

	fileBaseName := fileParts[0]
	fileHasEmoji := emoji.GetEmoji(fileBaseName) != ""

	// Если у файла уже есть эмодзи, и оно совпадает с текущим, пропускаем
	if fileHasEmoji && emoji.GetEmoji(fileBaseName) == inheritedEmoji {
		return nil
	}

	// Если у файла есть другое эмодзи, обновляем его на эмодзи родительской директории
	newName := fmt.Sprintf("%s %s %s", fileBaseName, inheritedEmoji, fileParts[1])
	newPath := filepath.Join(path, newName)

	if oldName != newName {
		if err := os.Rename(oldPath, newPath); err != nil {
			return fmt.Errorf("failed to rename file %s to %s: %w", oldPath, newPath, err)
		}
	}

	log.WithFields(log.Fields{
		"old_path": oldPath,
		"new_path": newPath,
	}).Info("Переименован файл")

	return nil
}
