package decorator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoratorDecorate_Full(t *testing.T) {
	// Создаём временную базовую директорию.
	baseDir := t.TempDir()

	// Создаём корневую директорию с эмодзи в имени, чтобы передать родительский эмодзи дальше.
	rootDir := filepath.Join(baseDir, "Root 😀")
	require.NoError(t, os.Mkdir(rootDir, 0755))

	// --- 1. Директория, которая будет переименована (так как не содержит эмодзи) ---
	// Создаём поддиректорию "SubDir" внутри корневой.
	subDir := filepath.Join(rootDir, "SubDir")
	require.NoError(t, os.Mkdir(subDir, 0755))

	// Файл внутри subDir, который должен получить эмодзи от родительской директории.
	file2 := filepath.Join(subDir, "2 something.md")
	require.NoError(t, os.WriteFile(file2, []byte("content"), 0644))

	// --- 2. Файл в корневой директории, который должен быть переименован ---
	// Имеет формат "BaseName rest", где BaseName не содержит эмодзи.
	file1 := filepath.Join(rootDir, "1 something.md")
	require.NoError(t, os.WriteFile(file1, []byte("content"), 0644))

	// --- 3. Файл, который не обрабатывается (не .md) ---
	fileOther := filepath.Join(rootDir, "other.txt")
	require.NoError(t, os.WriteFile(fileOther, []byte("content"), 0644))

	// --- 4. Файл, который уже содержит эмодзи – его имя не меняется ---
	fileAlready := filepath.Join(rootDir, "1 😀 decorated.md")
	require.NoError(t, os.WriteFile(fileAlready, []byte("content"), 0644))

	// --- 5. Директория, имя которой начинается с точки (skip-паттерн) ---
	hiddenDir := filepath.Join(rootDir, ".hidden")
	require.NoError(t, os.Mkdir(hiddenDir, 0755))
	fileHidden := filepath.Join(hiddenDir, "FileHidden something.md")
	require.NoError(t, os.WriteFile(fileHidden, []byte("content"), 0644))

	// --- 6. Директория, имя которой начинается с "_" (skip-паттерн) ---
	skipDir := filepath.Join(rootDir, "_skipDir")
	require.NoError(t, os.Mkdir(skipDir, 0755))
	fileSkip := filepath.Join(skipDir, "FileSkip something.md")
	require.NoError(t, os.WriteFile(fileSkip, []byte("content"), 0644))

	// --- 7. Директория, имя которой начинается с "any" (skip-паттерн) ---
	anyDir := filepath.Join(rootDir, "anyfolder")
	require.NoError(t, os.Mkdir(anyDir, 0755))
	fileAny := filepath.Join(anyDir, "FileAny something.md")
	require.NoError(t, os.WriteFile(fileAny, []byte("content"), 0644))

	// Создаём экземпляр декоратора с concurrencyLimit=2 и skipPatterns [".", "_", "any"]
	decor := NewDecorator(rootDir, 2, []string{".", "_", "any"})

	// Вызываем метод, который обходит и декорирует директории и файлы.
	err := decor.Decorate()
	require.NoError(t, err)

	// Проверяем, что корневая директория "😀Root" осталась без изменений.
	info, err := os.Stat(rootDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())

	// Файл "1 something.md" должен быть переименован в "1 😀 something.md"
	newFile1 := filepath.Join(rootDir, "1 😀 something.md")
	_, err = os.Stat(newFile1)
	require.NoError(t, err)

	// Файл "other.txt" остаётся без изменений.
	_, err = os.Stat(fileOther)
	require.NoError(t, err)

	// Файл "1 😀 decorated.md" остаётся без изменений.
	_, err = os.Stat(fileAlready)
	require.NoError(t, err)

	// Директория ".hidden" (и её содержимое) должна быть пропущена.
	info, err = os.Stat(hiddenDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())
	_, err = os.Stat(fileHidden)
	require.NoError(t, err)

	// Директория, имя которой начинается с "_" (skip), не переименовывается.
	info, err = os.Stat(skipDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())
	_, err = os.Stat(fileSkip)
	require.NoError(t, err)

	// Директория "anyfolder" не переименовывается (так как начинается с "any" по паттерну).
	info, err = os.Stat(anyDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())
	_, err = os.Stat(fileAny)
	require.NoError(t, err)

	// Проверяем, что поддиректория "SubDir" была переименована в "SubDir 😀"
	newSubDir := filepath.Join(rootDir, "SubDir 😀")
	info, err = os.Stat(newSubDir)
	require.NoError(t, err)
	require.True(t, info.IsDir())

	// Файл в поддиректории должен получить родительский эмодзи, т. е. переименоваться в "2 😀 something.md"
	newFile2 := filepath.Join(newSubDir, "2 😀 something.md")
	_, err = os.Stat(newFile2)
	require.NoError(t, err)
}

func TestDecorator_SkipDirectory(t *testing.T) {
	rootDir := t.TempDir()
	skipDirName := ".hiddenDir"
	skipDirPath := filepath.Join(rootDir, skipDirName)
	err := os.Mkdir(skipDirPath, 0755)
	require.NoError(t, err)

	fileName := "test something.md"
	filePath := filepath.Join(skipDirPath, fileName)
	err = os.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	decor := NewDecorator(rootDir, 2, []string{"."})
	err = decor.Decorate()
	require.NoError(t, err)

	// Директория и файл внутри должны остаться без изменений.
	_, err = os.Stat(skipDirPath)
	require.NoError(t, err)
	_, err = os.Stat(filePath)
	require.NoError(t, err)
}

func TestDecorator_RenameMarkdownFile(t *testing.T) {
	rootDir := t.TempDir()
	parentDir := filepath.Dir(rootDir)
	emojiRoot := filepath.Join(parentDir, "😀Root")
	err := os.Rename(rootDir, emojiRoot)
	require.NoError(t, err)
	rootDir = emojiRoot

	// Создаем поддиректорию без эмодзи в имени.
	subDirName := "SubDir"
	subDirPath := filepath.Join(rootDir, subDirName)
	err = os.Mkdir(subDirPath, 0755)
	require.NoError(t, err)

	fileName := "1 something.md"
	filePath := filepath.Join(subDirPath, fileName)
	err = os.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	decor := NewDecorator(rootDir, 2, []string{})
	err = decor.Decorate()
	require.NoError(t, err)

	expectedSubDirName := "SubDir 😀"
	expectedSubDirPath := filepath.Join(rootDir, expectedSubDirName)
	_, err = os.Stat(expectedSubDirPath)
	require.NoError(t, err)

	expectedFileName := "1 😀 something.md"
	expectedFilePath := filepath.Join(expectedSubDirPath, expectedFileName)
	_, err = os.Stat(expectedFilePath)
	require.NoError(t, err)
}

func TestDecorator_NonExistentDirectory(t *testing.T) {
	nonExistentDir := filepath.Join(t.TempDir(), "nonexistent")
	decor := NewDecorator(nonExistentDir, 2, []string{})
	// Функция Decorate не должна паниковать, даже если директория не существует.
	require.NotPanics(t, func() {
		err := decor.Decorate()
		require.NoError(t, err)
	})
}

func TestAddEmojiToFilename_NonMarkdown(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "test.txt"
	filePath := filepath.Join(tempDir, fileName)
	err := os.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	err = addEmojiToFilename(fileName, "😀", tempDir, filePath)
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	require.NoError(t, err)
}

func TestAddEmojiToFilename_NoSpaceInName(t *testing.T) {
	tempDir := t.TempDir()
	// Имя файла не содержит пробела для разделения, поэтому ничего не делаем.
	fileName := "test.md"
	filePath := filepath.Join(tempDir, fileName)
	err := os.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	err = addEmojiToFilename(fileName, "😀", tempDir, filePath)
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	require.NoError(t, err)
}

func TestAddEmojiToFilename_AlreadyHasEmoji(t *testing.T) {
	tempDir := t.TempDir()
	// Файл уже содержит эмодзи в базовой части имени.
	fileName := "1 😀 something.md"
	filePath := filepath.Join(tempDir, fileName)
	err := os.WriteFile(filePath, []byte("content"), 0644)
	require.NoError(t, err)

	err = addEmojiToFilename(fileName, "😀", tempDir, filePath)
	require.NoError(t, err)

	_, err = os.Stat(filePath)
	require.NoError(t, err)
}
