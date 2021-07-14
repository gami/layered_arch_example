package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

//go:embed env/default.toml
var defaultEnv []byte

//go:embed env/dev.toml
var devEnv []byte

//go:embed env/prd.toml
var prdEnv []byte

//Config represents configuration root.
type Config struct {
	Database Database
	AppEnv   string
}

var config *Config
var once sync.Once

//GetConfig is a function to get Configuration. This loading process occurs once in boot.
func GetConfig() *Config {
	once.Do(func() {
		cfg, err := NewConfig()
		if err != nil {
			panic(err)
		}
		config = cfg
	})
	return config
}

//NewConfig is a function to init and load Configuration from file or environment variables.
func NewConfig() (*Config, error) {
	conf := &Config{}
	conf.AppEnv = "dev"

	loadFromEnv(conf)

	err := loadFromToml(defaultEnv, conf)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to load env")
	}

	err = loadFromTomlEnv(conf)
	if err != nil {
		return nil, fmt.Errorf("fail loadFromToml by env : %v : env=%v", err, conf.AppEnv)
	}

	loadDatabaseConfig(conf)

	return conf, nil
}

func loadFromToml(tml []byte, conf *Config) error {
	_, err := toml.DecodeReader(bytes.NewBuffer(tml), conf)
	if err != nil {
		return fmt.Errorf("fail to decode toml : %v", err)
	}
	return nil
}

func loadFromTomlEnv(conf *Config) error {
	switch conf.AppEnv {
	case "dev":
		return loadFromToml(devEnv, conf)
	case "prd":
		return loadFromToml(prdEnv, conf)
	default:
		return nil
	}
}

func loadFromEnv(conf *Config) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "" {
		conf.AppEnv = appEnv
	}

	fmt.Printf("conf.AppEnv: %v\n", conf.AppEnv)
}
