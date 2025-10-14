package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env       string         `yaml:"env" env-default:"local" env:"ENV"`
	Database  DatabaseConfig `yaml:"database"`
	Port      int            `yaml:"port" env-default:"8080" env:"PORT"`
	JwtSecret string         `yaml:"jwtsecret" env:"JWT_SECRET"`
}

type DatabaseConfig struct {
	User     string `yaml:"user" env:"DB_USER"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
	Host     string `yaml:"host" env:"DB_HOST"`
	Name     string `yaml:"name" env:"DB_NAME"`
	Port     int    `yaml:"port" env:"DB_PORT"`
	SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	var cfg Config

	if path != "" {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			panic("config file does not exist: " + path)
		}
		fmt.Println("Loading config from file:", path)
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			panic("failed to read config: " + err.Error())
		}
	} else {
		fmt.Println("No config file found, loading from environment variables...")
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			panic("failed to read environment variables: " + err.Error())
		}
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
