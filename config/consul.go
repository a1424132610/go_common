package config

import (
	"github.com/asim/go-micro/v3/config"
	"github.com/go-micro/plugins/v3/config/source/consul"
	"strconv"
)

type ConsulConfig struct {
	//ip地址
	Host string
	Port int64
	//配置前缀
	Prefix string
}

//NewConsulConfig 创建一个consul配置对象
func NewConsulConfig(host string, port int64, prefix string) Strategy {
	return &ConsulConfig{
		Host:   host,
		Port:   port,
		Prefix: prefix,
	}
}

//GetConfig 获取到consul中的配置
func (c *ConsulConfig) GetConfig() (config.Config, error) {
	source := consul.NewSource(
		consul.WithAddress(c.Host+":"+strconv.FormatInt(c.Port, 10)),
		consul.WithPrefix(c.Prefix),
		consul.StripPrefix(true),
	)
	cfg, err := config.NewConfig()
	if err != nil {
		return cfg, err
	}
	err = cfg.Load(source)
	return cfg, err
}
