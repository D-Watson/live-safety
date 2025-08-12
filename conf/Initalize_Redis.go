package conf

import (
	"context"
	"os"

	"gopkg.in/yaml.v3"
	"live_safty/log"
)

var GlobalConfig *Config

type Config struct {
	DB *Database `yaml:"database"`
}
type Database struct {
	Address string `yaml:"address"`
	Passwd  string `yaml:"password"`
}

func ParseConfig(ctx context.Context) (*Config, error) {
	data, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Errorf(ctx, "[configs] read file error", err)
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Error(ctx, "解析配置失败: %v\n", err)
		return nil, err
	}
	return &config, nil
}
