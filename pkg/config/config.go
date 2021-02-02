package config

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
)

var Viper *viper.Viper

type Configuration map[string]interface{}

func init() {
	// 初始化
	Viper = viper.New()
	// 配置文件名称(无扩展名)
	Viper.SetConfigName(".env")
	// 如果配置文件的名称中没有扩展名，则需要配置此项
	Viper.SetConfigType("env")
	// 查找配置文件所在的路径
	Viper.AddConfigPath(".")
	// 查找读取配置文件并处理读取配置文件的错误
	if err := Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误
			panic(fmt.Errorf("Fatal error: %s \n", err))
		} else {
			// 配置文件被找到，但产生了另外的错误
			log.Fatalf("Read config failed. %v", err)
		}
	}
	// 设置环境变量前缀
	Viper.SetEnvPrefix("appenv")
	// viper.Get 请求时随时检查环境变量
	Viper.AutomaticEnv()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return Get(envName, defaultValue[0])
	}
	return Get(envName)
}

// Append 新增配置项
func Append(name string, configuration map[string]interface{}) {
	Viper.Set(name, configuration)
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(Get(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(Get(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(Get(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(Get(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(Get(path, defaultValue...))
}
