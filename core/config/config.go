package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	//xerror "github.com/pkg/errors"
)

func Parse(path string) error {
	// 从配置文件中读取配置
	if path != "" {
		viper.SetConfigFile(path)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		return errors.New("path can not be empty")
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		return  fmt.Errorf("Fatal error config file: %s \n", err)
	}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())

	return nil
}
