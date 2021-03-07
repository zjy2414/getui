package getui

// Config配置
type GeTuiConfig struct {
	AppId        string
	AppKey       string
	MasterSecret string

	RequestTimeout int64 // 单位：秒
}
