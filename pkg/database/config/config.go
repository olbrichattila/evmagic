package config

import (
	"fmt"
	"os"

	// This blank import is necessary to have the driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	EnvDBConnection = "DB_CONNECTION"
	envDBUserName   = "DB_USERNAME"
	envDBPassword   = "DB_PASSWORD"
	envDBHost       = "DB_HOST"
	envDBPort       = "DB_PORT"
	envDBDatabase   = "DB_DATABASE"
)

func GetConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		envWithDefault(envDBUserName, "user"),
		envWithDefault(envDBPassword, "mypassword"),
		envWithDefault(envDBHost, "localhost"),
		envWithDefault(envDBPort, "3306"),
		envWithDefault(envDBDatabase, "emagic"),
	)
}

func GetConnectionName() string {
	return "mysql"
}

func envWithDefault(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}
