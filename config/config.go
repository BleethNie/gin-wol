package config

var Cfg Config

// 配置文件的结构体
type Config struct {
	Server Server
	SQLite SQLite
}

type Server struct {
	AppMode   string
	BackPort  string
	FrontPort string
	UseHttps  bool
}

type SQLite struct {
	Name string
}
