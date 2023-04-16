package configs

import (
	"fmt"
	"log"
	"os"

	env "github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	NODE_ENV   string `env:"NODE_ENV"`
	Name       string `env:"NAME"`
	PORT       string `env:"PORT"`
	MONNGO_URL string `env:"DB_URL"`
	REDIS_URL  string `env:"REDIS_URL"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	fmt.Println("NODE_ENV", os.Getenv("NODE_ENV"))
	if os.Getenv("NODE_ENV") == "Development" {

		err := godotenv.Load()

		if err != nil {
			log.Fatal("Error loading .env file")
			return nil, err
		}

		err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
		if err != nil {
			log.Fatalf("unable to parse ennvironment variables: %e", err)
			return nil, err
		}
		return &cfg, nil
	}
	err := env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
		return nil, err
	}
	return &cfg, nil

}
