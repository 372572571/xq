package logx

type Config struct {
	Enable     bool
	Info       string // 其他日志
	Error      string // error 文件路径
	Panic      string // panic 文件路径
	Level      string // 输出等级  debug info warn
	// Skip       int    // zap.AddCallerSkip(1)
	// MaxAge     int    // 保存多少天
	// MaxBackups int    // 保存几份
}
