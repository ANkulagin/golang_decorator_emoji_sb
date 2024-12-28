package main

import (
	"flag"
	"github.com/ANkulagin/golang_decorator_emoji_sb/internal/config"
	"github.com/ANkulagin/golang_decorator_emoji_sb/internal/service/decorator"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "Путь к конфигурационному файлу")
	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Настройка уровня логирования
	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Не удалось установить уровень логирования: %v", err)
	}
	log.SetLevel(level)

	// Настройка формата логирования
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		// Можно добавить другие настройки
		ForceColors: true,
	})

	// Настройка вывода логов (можно перенаправить в файл, если нужно)
	log.SetOutput(os.Stdout)

	// Преобразование относительных путей в абсолютные
	absSrcDir, err := filepath.Abs(cfg.SrcDir)
	if err != nil {
		log.Fatalf("Не удалось определить абсолютный путь для исходной директории: %v", err)
	}

	log.Infof("Уровень логирования: %s", cfg.LogLevel)

	dec := decorator.NewDecorator(absSrcDir, cfg.ConcurrencyLimit)

	if err := dec.Decorate(); err != nil {
		log.Fatalf("Произошла ошибка при декорировании: %v", err)
	}

}
