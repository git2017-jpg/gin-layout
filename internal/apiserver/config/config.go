package config

import "github.com/BooeZhang/gin-layout/internal/apiserver/options"

// Config 配置
type Config struct {
	*options.Options
}

// CreateConfigFromOptions 基于给定的配置选项创建配置实例。
func CreateConfigFromOptions(opts *options.Options) (*Config, error) {
	return &Config{opts}, nil
}