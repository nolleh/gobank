package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var configPath = "."

func InitConfig(env string, config interface{}) error {
	SetConfigPath("config")
	return ReadConfig(env, config)
}

func SetConfigPath(in string) {
	configPath = in
	viper.AddConfigPath(configPath)
}

func ReadConfig(env string, config interface{}) error {
	viper.SetConfigName("default")
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	if env != "" {
		loadFileConfig(env + ".yaml")
	}
	loadFileConfig("local.yaml")

	if err := viper.Unmarshal(config); err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}

func loadFileConfig(filename string) error {
	f, err := os.Open(filepath.Join(configPath, filename))
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}

	viper.MergeConfig(f)
	return nil
}
