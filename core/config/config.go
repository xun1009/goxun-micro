package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"os"

	//xerror "github.com/pkg/errors"
)

func Parse(path string) error {
	// read from config file
	if path != "" {
		viper.SetConfigFile(path)   // file path name
		viper.SetConfigType("yaml") // file type
	} else {
		return errors.New("path can not be empty")
	}

	_,err := os.Stat(path)
	if err != nil{
		return 	fmt.Errorf("file not exist")
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		return  fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}
