package qlog

import (
	"go.uber.org/zap/zapcore"
)

type Conf struct {
	Debug    bool      `yaml:"debug" json:"debug"`
	LogCores []LogCore `yaml:"log_cores" json:"log_cores"`
}

type LogCore struct {
	LumberJack LumberJack `yaml:"lumber_jack" json:"lumber_jack"`
	Levels     Levels     `yaml:"levels" json:"levels"`
}

type LumberJack struct {
	Path        string `yaml:"path" json:"path"`                     //日志路径
	Name        string `yaml:"name" json:"name"`                     //日志文件名
	MaxSizeInMB int    `yaml:"max_size_in_mb" json:"max_size_in_mb"` //日志文件最大大小(MB)
	MaxBackups  int    `yaml:"max_backups" json:"max_backups"`       //日志文件最大保存数量
	MaxAgeInDay int    `yaml:"max_age_in_day" json:"max_age_in_day"` //日志文件最大保存时间(天)
	Compress    bool   `yaml:"compress" json:"compress"`             //是否压缩
}

type Levels []zapcore.Level

func (l Levels) Enabled(level2 zapcore.Level) bool {
	for _, level := range l {
		if level == level2 {
			return true
		}
	}
	return false
}
