package config

import (
	"encoding/json"
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

// LoadConfig 支持 yaml/json/env
func LoadConfig(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml"):
		return yaml.Unmarshal(data, v)
	case strings.HasSuffix(path, ".json"):
		return json.Unmarshal(data, v)
	case strings.HasSuffix(path, ".env"):
		return loadEnv(path, v)
	default:
		return errors.New("unsupported config file type")
	}
}

// 简单实现 .env 加载（只支持 key=value）
func loadEnv(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	envMap := make(map[string]string)
	bytes, _ := ioutil.ReadAll(file)
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) == 2 {
			envMap[kv[0]] = kv[1]
		}
	}

	// 把 map 转 JSON 再反序列化到 v
	js, _ := json.Marshal(envMap)
	return json.Unmarshal(js, v)
}
