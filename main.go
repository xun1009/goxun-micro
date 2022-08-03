package main

import (
	"fmt"
	"github.com/spf13/viper"
	"goxun-micro/core/config"
)

func main()  {
	config.Parse(".\\core\\config\\test.yaml")
	fmt.Println(viper.Get("Name"))
}