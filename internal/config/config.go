package config

import (
	"flag"
	"github.com/caarlos0/env/v10"
	"log"
)

type Config struct {
	ServerAddress        string `env:"RUN_ADDRESS"`
	DBDSN                string `env:"DATABASE_DSN"`
	AccrualSystemAddress string `env:"ACCRUAL_SYSTEM_ADDRESS"`
	MigrationDir         string
}

const (
	defaultServerAddress = "localhost:8080"
	defaultDBDSN         = "postgres://postgres:pwd@localhost:5432/postgres?sslmode=disable"
	defaultAccrualSystem = "localhost:8181"
	migrationDir         = "./internal/storage/migrations"
)

func NewConfig() *Config {
	return &Config{
		ServerAddress:        defaultServerAddress,
		DBDSN:                defaultDBDSN,
		AccrualSystemAddress: defaultAccrualSystem,
		MigrationDir:         migrationDir,
	}
}

func (c *Config) ParseFlags() error {
	flag.Func("a", "example: '-a localhost:8080'", func(addr string) error {
		c.ServerAddress = addr
		return nil
	})
	flag.Func("d", "example: '-d postgres://postgres:pwd@localhost:5432/postgres?sslmode=disable'", func(dsn string) error {
		c.DBDSN = dsn
		return nil
	})
	flag.Func("r", "example: '-r localhost:8181'", func(addr string) error {
		c.AccrualSystemAddress = addr
		return nil
	})
	flag.Parse()

	err := env.Parse(c)
	if err != nil {
		log.Println("Config is not available", err)
		return err
	}
	return nil
}
