package config

import (
	"os"
)

//Database represents configuration about database connection.
type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Name     string
}

func loadDatabaseConfig(conf *Config) {
	if conf.AppEnv == "test" {
		conf.Database.Host = os.Getenv("TEST_DB_HOST")
		conf.Database.User = os.Getenv("TEST_DB_USER")
		conf.Database.Password = os.Getenv("TEST_DB_PASSWORD")
		conf.Database.Port = 3306
		conf.Database.Name = "user_test"
	} else {
		conf.Database.Host = os.Getenv("USER_DB_HOST")
		conf.Database.User = "user"
		conf.Database.Password = os.Getenv("USER_DB_PASSWORD")
		conf.Database.Port = 3306
		conf.Database.Name = "user"
	}
}
