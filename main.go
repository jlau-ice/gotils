package main

import (
	"fmt"

	"github.com/jlau-ice/gotils/config"
	"github.com/jlau-ice/gotils/log"
	"github.com/jlau-ice/gotils/str"
)

type AppConfig struct {
	Golang struct {
		Application struct {
			Name string `yaml:"name"`
		} `yaml:"application"`
	} `yaml:"golang"`
}

func main() {
	log.Info("Starting gotils demo...")
	// 测试字符串工具
	fmt.Println("CamelToSnake:", str.CamelToSnake("HelloWorld"))
	fmt.Println("SnakeToCamel:", str.SnakeToCamel("hello_world"))
	fmt.Println("ReverseStr:", str.ReverseStr("abcdef"))
	fmt.Println("RandomString:", str.RandomString(8))
	// 测试配置加载
	cfg := AppConfig{}
	err := config.LoadConfig("test.yaml", &cfg)
	if err != nil {
		log.Error("Load config failed: " + err.Error())
	} else {
		log.Info(fmt.Sprintf("Config: %+v", cfg))
	}

	log.Info("Demo finished.")
}
