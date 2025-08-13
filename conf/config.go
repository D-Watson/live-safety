package conf

import (
	"context"
	"os"

	"gopkg.in/yaml.v3"
	"live_safty/log"
)

var GlobalConfig *Config

type Config struct {
	DB     *Database `yaml:"database"`
	Server *Server   `yaml:"server"`
}

type Server struct {
	Name string `yaml:"name"`
	Rpc  *Rpc   `yaml:"rpc"`
	Http *Http  `yaml:"http"`
}

type Rpc struct {
	Host string `yaml:"host"`
}
type Http struct {
	Host string `yaml:"host"`
}

type Database struct {
	Address string `yaml:"address"`
	Passwd  string `yaml:"password"`
}

func ParseConfig(ctx context.Context) error {
	data, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Errorf(ctx, "[configs] read file error", err)
		return err
	}
	if err = yaml.Unmarshal(data, &GlobalConfig); err != nil {
		log.Error(ctx, "解析配置失败: %v\n", err)
		return err
	}
	return nil
}
