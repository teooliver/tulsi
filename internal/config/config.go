package config

type Config struct {
	Postgres PostgresConfig `envconfig:"POSTGRES"`
}

type PostgresConfig struct {
	DSN string `envconfig:"DSN"`
}
