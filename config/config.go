package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var Conf Config

// Config 总配置结构体
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Etcd   EtcdConfig   `mapstructure:"etcd"`
	Jaeger JaegerConfig `mapstructure:"jaeger"`
}

// ServerConfig 服务端配置 (对应 server key)
type ServerConfig struct {
	User ServiceInfo `mapstructure:"user"`
	Todo ServiceInfo `mapstructure:"todo"`
}

// ServiceInfo 具体服务的配置 (复用于 User 和 Todo)
type ServiceInfo struct {
	Name    string   `mapstructure:"name"`
	Address []string `mapstructure:"address"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
	Charset  string `mapstructure:"charset"`
}

// EtcdConfig Etcd 配置
type EtcdConfig struct {
	Endpoints []string `mapstructure:"endpoints"`
}

// JaegerConfig 链路追踪配置
type JaegerConfig struct {
	Address string `mapstructure:"address"`
}

// Init 初始化配置
func Init() {
	// 获取当前工作目录
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %s", err)
	}

	// 设置配置文件名
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("yaml")
	// 设置查找路径
	viper.AddConfigPath(filepath.Join(workDir, "config")) 
	viper.AddConfigPath(workDir)      
	viper.AddConfigPath(filepath.Join(workDir, "../../config"))                    

	// 读取配置
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 映射到结构体
	if err := viper.Unmarshal(&Conf); err != nil {
		log.Fatalf("Unable to decode into struct: %s", err)
	}

	log.Println("Config loaded successfully!")
}