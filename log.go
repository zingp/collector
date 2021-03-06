package main

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

func initLogs(filename string) (err error) {
	config := make(map[string]interface{})
	config["filename"] = filename
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}