package options

// CasBinOptions casBin配置
type CasBinOptions struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}

// NewCasBinOptions 创建默认的casBin配置项
func NewCasBinOptions() *CasBinOptions {
	return &CasBinOptions{
		ModelPath: "",
	}
}
