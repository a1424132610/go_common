package config

import "github.com/asim/go-micro/v3/config"

type Strategy interface {

	//GetConfig 获取到配置信息
	GetConfig() (config.Config, error)
}
