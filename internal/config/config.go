package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Env    string       `yaml:"env" env-default:"local"`
	HTTP   HTTPConfig   `yaml:"http"`
	Client ClientConfig `yaml:"client"`
}

type HTTPConfig struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"read-timeout"`
	WriteTimeout time.Duration `yaml:"write-timeout"`
}

type ClientConfig struct {
	Url string `yaml:"url"`
}

func MustLoad() *Config {

	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	return MustLoadPath(path)
}

func MustLoadPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	res = os.Getenv("CONFIG_PATH")

	if res == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("failed to load .env file: " + err.Error())
		}
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
