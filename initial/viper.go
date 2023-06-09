package initial

import (
	appConfig "backend/config"
	"backend/global"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper
// 优先级: 命令行 > 环境变量 > 默认值
func Viper() *viper.Viper {
	var config string

	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config == "" { // 判断命令行参数是否为空
		if configEnv := os.Getenv(appConfig.EnvConfigFile); configEnv == "" {
			config = appConfig.DefaultConfigFile
			fmt.Println("use default config file")
		} else {
			config = configEnv
			fmt.Println("use config file from env", appConfig.EnvConfigFile)
		}
	} else {
		fmt.Println("use config file", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.AppConfig); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.AppConfig); err != nil {
		fmt.Printf("load app config %s failed\n", config)
		fmt.Println(err)
	} else {
		fmt.Printf("load app config: %s \n", config)
		fmt.Println(global.AppConfig)
	}

	return v
}
