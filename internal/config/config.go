package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/redis/go-redis/v9"
)

type (
	Database struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Password string `yaml:"password"`
		Name     int    `yaml:"name" env-default:"0"`
	}

	Server struct {
		Host string `yaml:"host" env-required:"true"`
		Port string `yaml:"port" env-required:"true"`
	}

	Config struct {
		Database Database `yaml:"database"`
		Server   Server   `yaml:"server"`
		Env      string   `yaml:"env" env-default:"local"`
	}
)

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath)
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); err != nil {
		panic("config file does not exits: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config file: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "../config/config.yml", "config file path")

	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath
}

func NewRedis(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Database.Host + ":" + cfg.Database.Port,
		Password: cfg.Database.Password,
		DB:       cfg.Database.Name,
	})
}
