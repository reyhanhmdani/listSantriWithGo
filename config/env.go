package config

type Config struct {
	DBUsername string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASS"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     int    `envconfig:"DB_PORT"`
	DBName     string `envconfig:"DB_NAME"`
}
