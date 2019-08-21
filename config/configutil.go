package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var configPath = "."

// config
type Config struct {
	Database struct {
		Driver     string
		Connection string
	}
	HttpPort string

	DataStore struct {
		ProjectId string
		KeyFile string
	}
}

var config Config

func GetConfig(env string, path string) (*Config, error) {
	setConfigPath(path)
	err := readConfig(env, &config)
	return &config, err
}

func setConfigPath(in string) {
	configPath = in
	viper.AddConfigPath(configPath)
}

func readConfig(env string, config interface{}) error {
	viper.SetConfigName("default")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if env != "" {
		if err := loadFileConfig(env + ".yaml"); err != nil {
			return fmt.Errorf("Fatal error config file: %s \n", err)
		}
	}
	if err := loadFileConfig("local.yaml"); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}

func loadFileConfig(filename string) error {
	f, err := os.Open(filepath.Join(configPath, filename))
	if f != nil {
		defer func() {
			err := f.Close()
			fmt.Println(fmt.Errorf("Fatal error while close file: %s \n", err))
		}()
	}
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	return viper.MergeConfig(f)
}