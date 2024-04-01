package qconfig

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadFromFlags(configFileExt string, configPointer any) error {
	var confPath string
	flag.StringVar(&confPath, "c", "dev.yaml", "配置文件路径")
	flag.Parse()
	if confPath == "" {
		return fmt.Errorf("配置文件路径为空")
	}
	tomlData, err := os.ReadFile(confPath)
	if err != nil {
		return err
	}
	switch configFileExt {
	case ".yaml":
		return yaml.Unmarshal(tomlData, configPointer)
	case ".toml":
		return toml.Unmarshal(tomlData, configPointer)
	case ".json":
		return json.Unmarshal(tomlData, configPointer)
	default:
		return fmt.Errorf("不支持的配置文件类型: %s", configFileExt)
	}
}
