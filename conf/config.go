package conf

import (
	"context"
	"os"
	"strings"

	"github.com/D-Watson/live-safety/log"
	"gopkg.in/yaml.v3"
)

var GlobalConfig *Config

type Config struct {
	DB     *Database `yaml:"databases"`
	Server *Server   `yaml:"server"`
	Kafka  *Kafka    `yaml:"kafka"`
}

type FormatedConfig struct {
	DB     *Database  `yaml:"databases"`
	Server *Server    `yaml:"server"`
	Kafka  *KafkaConf `yaml:"kafka"`
}

type Server struct {
	Name    string `yaml:"name"`
	Rpc     *Rpc   `yaml:"rpc"`
	Http    *Http  `yaml:"http"`
	Timeout int    `yaml:"timeout"`
}

type Rpc struct {
	Host       string `yaml:"host"`
	ServerHost string `yaml:"serverHost"`
}
type Http struct {
	Host string `yaml:"host"`
}

type Database struct {
	Redis *RedisConf `yaml:"redis"`
	Mysql *MysqlConf `yaml:"mysql"`
}
type MysqlConf struct {
	UserName string  `yaml:"username"`
	Password string  `yaml:"password"`
	Address  string  `yaml:"address"`
	DBName   string  `yaml:"dbname"`
	Options  Options `yaml:"options"`
}
type Options struct {
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`
	Timeout      int `yaml:"timeout"`
	ReadTimeout  int `yaml:"readtimeout"`
	WriteTimeout int `yaml:"writetimeout"`
}
type RedisConf struct {
	Address string `yaml:"address"`
	Passwd  string `yaml:"password"`
}

type Kafka struct {
	Address  []string
	Topic    string
	MinBytes int
	MaxBytes int
}

type KafkaConf struct {
	Address  string `yaml:"address"`
	Topic    string `yaml:"topic"`
	MinBytes int    `yaml:"min_bytes"`
	MaxBytes int    `yaml:"max_bytes"`
}

func ParseConfig(ctx context.Context) error {
	data, err := os.ReadFile("./conf/config.yaml")
	if err != nil {
		log.Errorf(ctx, "[configs] read file error", err)
		return err
	}
	config := &FormatedConfig{}
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Error(ctx, "解析配置失败: %v\n", err)
		return err
	}
	conf := &Config{
		DB:     config.DB,
		Server: config.Server,
	}
	if config.Kafka != nil {
		address := strings.Split(config.Kafka.Address, ",")
		conf.Kafka = &Kafka{
			Address:  address,
			Topic:    config.Kafka.Topic,
			MinBytes: config.Kafka.MinBytes,
			MaxBytes: config.Kafka.MaxBytes,
		}
	}
	GlobalConfig = conf
	return nil
}
