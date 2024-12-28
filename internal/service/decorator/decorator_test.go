package decorator

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestDecorateDirectories_Success(t *testing.T) {
	path := t.TempDir()
	sut := NewDecorator(path)

	testCases := []struct {
		name  string
		setup func() (path, srcDir, destDir string, cleanup func())
	}{
		{
			name: "Успешное чтение файла без FrontMatter",
			setup: func() (path, srcDir, destDir string, cleanup func()) {
				srcDir, _ = os.MkdirTemp("", "src_dir")
				destDir, _ = os.MkdirTemp("", "dest_dir")
				path = filepath.Join(srcDir, "test_ok.md")
				f, _ := os.Create(path)
				_ = f.Close()

				return path, srcDir, destDir, func() {
					_ = os.RemoveAll(srcDir)
					_ = os.RemoveAll(destDir)
				}
			},
		},
		{
			name: "Успешное чтение файла с корректным FrontMatter",
			setup: func() (path, srcDir, destDir string, cleanup func()) {
				srcDir, _ = os.MkdirTemp("", "src_dir")
				destDir, _ = os.MkdirTemp("", "dest_dir")
				path = filepath.Join(srcDir, "test_ok_fm.md")
				content := `---
date: 2024-12-09
author: "ANkulagin"
tags:
  - "#daily"
  - "#notes"
closed: false
---
# Заголовок

Контент...
`
				err := os.WriteFile(path, []byte(content), 0644)
				require.NoError(t, err)

				return path, srcDir, destDir, func() {
					_ = os.RemoveAll(srcDir)
					_ = os.RemoveAll(destDir)
				}
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			path, srcDir, destDir, cleanup := tc.setup()
			defer cleanup()

			err := sut.ConvertFile(path, srcDir, destDir)

			require.NoError(t, err)
		})
	}
}
