package qconfig

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func Load(confPath string, configPointer any) error {
	if confPath == "" {
		return fmt.Errorf("配置文件路径为空")
	}
	tomlData, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	ext := filepath.Ext(confPath)
	switch ext {
	case ".yaml":
		return yaml.Unmarshal(tomlData, configPointer)
	case ".toml":
		return toml.Unmarshal(tomlData, configPointer)
	case ".json":
		return json.Unmarshal(tomlData, configPointer)
	default:
		return fmt.Errorf("不支持的配置文件类型: %s", ext)
	}
}

func LoadConfig(configModel any) error {
	var confPath string
	flag.StringVar(&confPath, "c", "devConf.toml", "toml配置文件路径")
	flag.Parse()
	if confPath == "" {
		return fmt.Errorf("toml配置文件路径为空")
	}
	tomlData, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	return toml.Unmarshal(tomlData, configModel)
}

func LoadToml(configModel any) error {
	var confPath string
	flag.StringVar(&confPath, "c", "devConf.toml", "toml配置文件路径")
	flag.Parse()
	if confPath == "" {
		return fmt.Errorf("toml配置文件路径为空")
	}
	tomlData, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	return toml.Unmarshal(tomlData, configModel)
}

func LoadYaml(configModel any) error {
	var confPath string
	flag.StringVar(&confPath, "c", "devConf.toml", "toml配置文件路径")
	flag.Parse()
	if confPath == "" {
		return fmt.Errorf("toml配置文件路径为空")
	}
	tomlData, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(tomlData, configModel)
}
