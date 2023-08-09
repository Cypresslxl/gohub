package config

import (
	"gohub/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

// viper 库实例
var viper *viperlib.Viper

// Configue 动态加载配置信息
type ConfigFunc func() map[string]interface{}

// 先加载到此数组，loadConfig 再动态生成配置信息
var ConfigFuncs map[string]ConfigFunc

func init() {
	// 1.初始化Viper库
	viper = viperlib.New()
	// 2.配置类型支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("env")

	// 3.环境变量配置文件查找的路径，相对于 main.go文件的路径
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// InitConfig 初始化配置信息，完成对环境变量以及 config 信息的加载
func InitConfig(env string) {
	// 1. 加载环境变量
	loadEnv(env)
	// 2. 注册配置信息
	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn()) //因为fn确实是函数引用
	}
}

func loadEnv(envSuffix string) {
	// 默认加载 .env 文件，如果有传参 --env=name 的话，加载 .env.name 文件
	envPath := ".env"
	if len(envSuffix) > 0 {
		filename := envPath + envSuffix
		if _, err := os.Stat(filename); err == nil { //os.Stat() 的作用：It returns an os.FileInfo struct that contains various information about the file or directory.
			// 如 .env.testing 或 .env.stage
			envPath = filename
		}
	}

	// 加载 env
	viper.SetConfigName(envPath) //SetConfigName sets name for the config file.
	if err := viper.ReadInConfig(); err != nil {
		//ReadInConfig will discover and load the configuration file from disk
		// and key/value stores, searching in one of the defined paths.
		panic(err)
	}

	// 监控 .env 文件，变更时重新加载
	viper.WatchConfig()
}

// Env 读取环境变量，支持默认值
func Env(envName string, defaultValue ...interface{}) interface{} {
	if len(defaultValue) > 0 {
		return internalGet(envName, defaultValue[0])
	}
	return internalGet(envName)
}

// Add 新增配置项
func Add(name string, configFn ConfigFunc) {
	ConfigFuncs[name] = configFn
}

// Get 获取配置项
// 第一个参数 path 允许使用点式获取，如：app.name
// 第二个参数允许传参默认值
func Get(path string, defaultValue ...interface{}) string {
	return GetString(path, defaultValue...)
}

// In summary, this function is designed to retrieve configuration values using Viper,
// with the option to provide a default value if the requested value is missing or empty 1
// If no default value is provided, the function returns nil in such cases.
func internalGet(path string, defaultValue ...interface{}) interface{} {
	// Check if the configuration key exists or if its value is empty
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		// If a default value is provided and there is at least one element in defaultValue slice
		if len(defaultValue) > 0 {
			return defaultValue[0] // Return the provided default value
		}
		return nil // Return nil since there's no default value
	}
	// Return the value from the configuration
	return viper.Get(path)
}

// next
// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
	return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
	return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取Float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
	return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(internalGet(path, defaultValue...))
}

// // 用泛型实现上面的代码,这里会出现interface {} is string, not int，
// 因为viper.Get 返回的是 sting 类型数据，如果.env 设置了配置需要取 int 类型，按这种写法 Get [int] 的话就会 interface {} is string, not int
// func Get[T any](path string, defaultValue ...interface{}) T {
// 	if value := internalGet(path, defaultValue...); value != nil {
// 		return value.(T) //类型转换
// 	}
// 	// 泛型不能返回 nil，因此需要根据类型建立空变量，这样返回的会是对应类型的"空"值
// 	var tmp T
// 	return tmp
// }

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
