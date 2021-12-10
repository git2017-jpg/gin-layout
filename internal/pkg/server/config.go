package server

import (
	"github.com/BooeZhang/gin-layout/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net"
	"strconv"
	"strings"
	"time"
)

const (
	// RecommendedEnvPrefix 定义所有服务使用的 ENV 前缀.
	RecommendedEnvPrefix = "API-SERVER"
)

// Config API 服务配置结构体
type Config struct {
	RunInfo   *RunInfo
	Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

// CertKey 证书相关
type CertKey struct {
	// CertFile PEM 编码证书的文件
	CertFile string
	// KeyFile PEM 编码私钥的文件
	KeyFile string
}

// RunInfo 服务器运行配置。
type RunInfo struct {
	BindAddress string
	BindPort    int
	CertKey     CertKey
}

// Address 将主机 IP 地址和主机端口号连接成一个地址字符串，例如：0.0.0.0:8443。
func (r *RunInfo) Address() string {
	return net.JoinHostPort(r.BindAddress, strconv.Itoa(r.BindPort))
}

// JwtInfo JWT相关配置
type JwtInfo struct {
	// defaults to "jwt"
	Realm string
	// defaults to empty
	Key string
	// defaults to one hour
	Timeout time.Duration
	// defaults to zero
	MaxRefresh time.Duration
}

// NewConfig 返回具有默认值的 Config 结构。
func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,
		Jwt: &JwtInfo{
			Realm:      "jwt",
			Timeout:    1 * time.Hour,
			MaxRefresh: 1 * time.Hour,
		},
	}
}

// CompletedConfig 完整配置信息
type CompletedConfig struct {
	*Config
}

// Complete 填写任何未设置的字段
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}

// New 从给定的配置返回 GenericAPIServer 的新实例。
func (c CompletedConfig) New() (*APIServer, error) {
	s := &APIServer{
		RunInfo: c.RunInfo,
		Mode:              c.Mode,
		Healthz:           c.Healthz,
		EnableMetrics:     c.EnableMetrics,
		EnableProfiling:   c.EnableProfiling,
		Middlewares:       c.Middlewares,
		Engine:            gin.New(),
	}

	InitGenericAPIServer(s)

	return s, nil
}

// LoadConfig reads in config file and ENV variables if set.
func LoadConfig(cfg string, defaultName string) {
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc")
		viper.SetConfigName(defaultName)
	}

	// Use config file from the flag.
	viper.SetConfigType("yaml")              // set the type of the configuration to yaml.
	viper.AutomaticEnv()                     // read in environment variables that match.
	viper.SetEnvPrefix(RecommendedEnvPrefix) // set ENVIRONMENT variables prefix to IAM.
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("WARNING: viper failed to discover and load the configuration file: %s", err.Error())
	}
}
