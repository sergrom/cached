package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Db struct {
		Host string `yaml:"host"`
		Socket string `yaml:"socket"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Proto string `yaml:"proto"`
		Address string `yaml:"address"`
		Database string `yaml:"database"`
	}
	Server struct{
		InvalidatePeriod int64 `yaml:"invalidate_cache_period"`
		JsonRpcSocket string `yaml:"json_rpc_socket"`
	}
}

func (cfg *Config) LoadConfig(cfgFile string) error {
	cfgContent, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return err
	}
	c := &Config{}
	err = yaml.Unmarshal(cfgContent, c)
	if err != nil {
		return err
	}

	errorMsg := c.validateConfig()
	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	*cfg = *c

	return nil
}

func (cfg *Config) validateConfig() string {
	// Todo
	return ""
}