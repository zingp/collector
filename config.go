package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

type AppConf struct {
	listenFile string
	log string
}

var appConf AppConf

func initConfig(file string) (err error){
	conf, err := config.NewConfig("ini", file)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}
	appConf.listenFile = conf.String("listenFile")
	appConf.log = conf.String("log")
	return
}