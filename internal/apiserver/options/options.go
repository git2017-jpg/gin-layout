package options

import (
	"encoding/json"
	options2 "github.com/BooeZhang/gin-layout/internal/pkg/options"
	"github.com/BooeZhang/gin-layout/pkg/log"
)

var o *Options

// Options 配置选项
type Options struct {
	ServerRunOptions  *options2.ServerRunOptions  `json:"server"   mapstructure:"server"`
	GRPCOptions       *options2.GRPCOptions       `json:"grpc"     mapstructure:"grpc"`
	HttpServerOptions *options2.HttpServerOptions `json:"http"     mapstructure:"http"`
	MySQLOptions      *options2.MySQLOptions      `json:"mysql"    mapstructure:"mysql"`
	RedisOptions      *options2.RedisOptions      `json:"redis"    mapstructure:"redis"`
	JwtOptions        *options2.JwtOptions        `json:"jwt"      mapstructure:"jwt"`
	Log               *log.Options                `json:"log"      mapstructure:"log"`
	FeatureOptions    *options2.FeatureOptions    `json:"feature"  mapstructure:"feature"`
}

// NewOptions 生成配置选项
func NewOptions() *Options {
	o = &Options{
		ServerRunOptions:  options2.NewServerRunOptions(),
		GRPCOptions:       options2.NewGRPCOptions(),
		HttpServerOptions: options2.NewHttpServerOptions(),
		MySQLOptions:      options2.NewMySQLOptions(),
		RedisOptions:      options2.NewRedisOptions(),
		JwtOptions:        options2.NewJwtOptions(),
		Log:               log.NewOptions(),
		FeatureOptions:    options2.NewFeatureOptions(),
	}

	return o
}

// String 配置字符输出
func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

func GetOptions() *Options {
	return o
}