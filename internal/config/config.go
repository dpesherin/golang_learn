package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Port       string `yaml:"port" env:"PORT" env-default:"5432"`
	Host       string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name       string `yaml:"name" env:"NAME" env-default:"postgres"`
	User       string `yaml:"user" env:"USER" env-default:"user"`
	Password   string `yaml:"password" env:"PASSWORD"`
	HttpServer `yaml:"httpServer"`
}

type HttpServer struct {
	Address string `yaml:"address" env-default:"0.0.0.0:8080"`
}

func MustLoad() *Config {
	// Получаем путь до конфиг-файла из env-переменной CONFIG_PATH
	configPath := "./config/local.yaml"

	// Проверяем существование конфиг-файла
	if _, err := os.Stat(configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	var cfg Config

	// Читаем конфиг-файл и заполняем нашу структуру
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("error reading config file: %s", err)
	}
	fmt.Println(cfg)

	return &cfg
}
