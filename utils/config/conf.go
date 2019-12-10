package utils

import (
	"github.com/go-ini/ini"
	"path/filepath"
)

type Config struct {
	conf *ini.File
}

var (
	c = new(Config)
	//项目根路径
	appPath string
	//配置文件路径
	confPath string
	//当前选择的配置文件
	RunMode  string
	HttpPort string
)

func init() {
	c.conf = parseConfig()
	// 只读操作增加性能
	c.conf.BlockMode = false
	RunMode = c.conf.Section("").Key("runmode").String()
	HttpPort = ":" + c.conf.Section("").Key("httpport").String()
}

func String(key string) string {
	return c.conf.Section("").Key(key).String()
}

func Int(key string) int {
	r, err := c.conf.Section("").Key(key).Int()
	if err != nil {
		return 0
	}
	return r
}

func Bool(key string) bool {
	r, err := c.conf.Section("").Key(key).Bool()
	if err != nil {
		return false
	}
	return r
}

func parseConfig() *ini.File {
	// 载入入口配置文件
	indexConfPath := filepath.Join(confPath, "./conf/app.ini")
	conf, err := ini.Load(indexConfPath)
	if err != nil {
		panic(err)
	}
	return conf
}
