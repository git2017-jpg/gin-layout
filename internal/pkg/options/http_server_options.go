package options

import "github.com/BooeZhang/gin-layout/internal/pkg/server"

// HttpServerOptions http服务配置项.
type HttpServerOptions struct {
	BindAddress string `json:"bind-address" mapstructure:"bind-address"`
	// BindPort 设置 Listener 时被忽略，即使为 0 也会提供 HTTPS。
	BindPort int `json:"bind-port"    mapstructure:"bind-port"`
	// Required 设置为 true 意味着 BindPort 不能为零。
	Required bool
	// ServerCert TLS 证书信息
	ServerCert GeneratableKeyCert `json:"tls"          mapstructure:"tls"`
}

// CertKey 证书相关配置
type CertKey struct {
	CertFile string `json:"cert-file"        mapstructure:"cert-file"`
	KeyFile  string `json:"private-key-file" mapstructure:"private-key-file"`
}

// GeneratableKeyCert 生成证书的相关配置.
type GeneratableKeyCert struct {
	CertKey       CertKey `json:"cert-key" mapstructure:"cert-key"`
	CertDirectory string  `json:"cert-dir"  mapstructure:"cert-dir"`
	PairName      string  `json:"pair-name" mapstructure:"pair-name"`
}

// NewHttpServerOptions 创建一个带有默认参数的http服务.
func NewHttpServerOptions() *HttpServerOptions {
	return &HttpServerOptions{
		BindAddress: "0.0.0.0",
		BindPort:    8443,
		Required:    true,
		ServerCert: GeneratableKeyCert{
			PairName:      "key-cert",
			CertDirectory: "/var/run/key-cert",
		},
	}
}

// ApplyTo 赋值配置项
func (h *HttpServerOptions) ApplyTo(c *server.Config) error {
	// SecureServing is required to serve https
	c.RunInfo = &server.RunInfo{
		BindAddress: h.BindAddress,
		BindPort:    h.BindPort,
		CertKey: server.CertKey{
			CertFile: h.ServerCert.CertKey.CertFile,
			KeyFile:  h.ServerCert.CertKey.KeyFile,
		},
	}

	return nil
}