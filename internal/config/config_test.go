package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig_Success(t *testing.T) {
	content := `
src_dir: "/some/path"
log_level: "debug"
concurrency_limit: 5
skip_patterns:
  - "."
  - "_"
`
	tmpFile, err := os.CreateTemp("", "config_*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	cfg, err := LoadConfig(tmpFile.Name())
	require.NoError(t, err)
	require.Equal(t, "/some/path", cfg.SrcDir)
	require.Equal(t, "debug", cfg.LogLevel)
	require.Equal(t, 5, cfg.ConcurrencyLimit)
	require.ElementsMatch(t, []string{".", "_"}, cfg.SkipPatterns)
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	_, err := LoadConfig("nonexistent.yaml")
	require.Error(t, err)
	require.Contains(t, err.Error(), "ошибка чтения конфигурационного файла")
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	content := `
src_dir: "/some/path"
log_level: "debug"
concurrency_limit: not_a_number
`
	tmpFile, err := os.CreateTemp("", "config_invalid_*.yaml")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	_, err = LoadConfig(tmpFile.Name())
	require.Error(t, err)
	require.Contains(t, err.Error(), "ошибка разбора конфигурации")
}
