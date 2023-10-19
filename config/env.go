package config

type Config struct {
	DBUsername string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASS"`
	DBHost     string `envconfig:"DB_HOST"`
	DBPort     int    `envconfig:"DB_PORT"`
	DBName     string `envconfig:"DB_NAME"`

	//SMPT
	EmailFrom string `envconfig:"EMAIL_FROM"`
	SMTPHost  string `envconfig:"SMTP_HOST"`
	SMTPPort  int    `envconfig:"SMTP_PORT"`
	SMTPUser  string `envconfig:"SMTP_USER"`
	SMTPPass  string `envconfig:"SMTP_PASS"`
}
