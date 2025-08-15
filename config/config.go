package config

func NewConfig() Config {
	return Config{
		Server:   ServerConfig{},
		Logger:   LoggerConfig{},
		Postgres: PostgresConfig{},
	}
}

type Config struct {
	Server   ServerConfig
	Logger   LoggerConfig
	Postgres PostgresConfig
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

type LoggerConfig struct {
	Level    string `env:"LOG_LEVEL" envDefault:"info"`
	Encoding string `env:"LOG_ENCODING" envDefault:"json"`
}

type PostgresConfig struct {
	Database string `env:"POSTGRES_DATABASE" envDefault:"microservice"`
	Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
	User     string `env:"POSTGRES_USER" envDefault:"postgres"`
	Password string `env:"POSTGRES_PASSWORD" envDefault:"843904831171"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	SSLMode  string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
}
