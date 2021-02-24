package config

import (
	"fmt"
	"github.com/marshhu/ma-frame/utils"
	"github.com/spf13/viper"
	"log"
)

type ConfigSetting struct {
	CfgFile string   // 配置文件名
	CfgDirs []string // 配置路径
	CfgType string   // 配置文件格式
}

func Init(config ConfigSetting) {
	viper.SetConfigName(config.CfgFile) // name of config file (without extension)
	viper.SetConfigType(config.CfgType) // REQUIRED if the config file does not have the extension in the name
	for _, dir := range config.CfgDirs {
		if utils.IsEmpty(dir) {
			continue
		}
		viper.AddConfigPath(dir) // path to look for the config file in
	}

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("加载配置文件失败: %s \n", err))
	}
	log.Println("使用的配置文件位置：" + viper.ConfigFileUsed())
}
