package main

import (
	"flag"
	"github.com/ANkulagin/golang_decorator_emoji_sb/internal/config"
	"log"
)

func main() {
	configPath := flag.String("config", "configs/config.yaml", "Путь к конфигурационному файлу")
	flag.Parse()
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}

	// Преобразование относительных путей в абсолютные
	//absSrcDir, err := filepath.Abs(cfg.SrcDir)
	//if err != nil {
	//	log.Fatalf("Не удалось определить абсолютный путь для исходной директории: %v", err)
	//}

	log.Printf("Уровень логирования: %s\n", cfg.LogLevel)
}
