package config
import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
    Env string `yaml:"env"`
}


func MustLoad() *Config {
	configPath := os.Getenv("BASE_CONFIG_PATH")
	if configPath == "" {
		panic("BASE_CONFIG_PATH environment variable is not set")
	}
	if _, err := os.Stat(configPath); err != nil {
		panic("error opening config file: %s")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		panic("error reading config file: %s")
	}

	return &cfg
}

func InitEnv() {
	op := "pkg.InitEnv "
	err := godotenv.Load()
	if err != nil {
		panic(op + "Error loading .env file")
	}
}


