package qconfig

import (
	"encoding/json"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type testModel struct {
	Value1 int
	Value2 string
	value3 int
	Value4 struct {
		Value41 int
		Value42 string
		value43 int
	}
}

func TestLoad(t *testing.T) {
	m := testModel{
		Value1: 1,
		Value2: "2",
		Value4: struct {
			Value41 int
			Value42 string
			value43 int
		}{Value41: 41, Value42: "42"},
	}
	yamlData, err := yaml.Marshal(&m)
	if err != nil {
		t.Log(err)
		return
	}
	if err := os.WriteFile("yamlData.yaml", yamlData, 0666); err != nil {
		t.Log(err)
		return
	}

	tomlData, err := toml.Marshal(&m)
	if err != nil {
		t.Log(err)
		return
	}
	if err := os.WriteFile("tomlData.toml", tomlData, 0666); err != nil {
		t.Log(err)
		return
	}

	jsonData, err := json.Marshal(&m)
	if err != nil {
		t.Log(err)
		return
	}
	if err := os.WriteFile("jsonData.json", jsonData, 0666); err != nil {
		t.Log(err)
		return
	}

	for _, path := range []string{"yamlData.yaml", "tomlData.toml", "jsonData.json"} {
		m := new(testModel)
		if err := Load(path, m); err != nil {
			t.Log(err)
			return
		}
		t.Logf("%+v\n", m)
	}
}
