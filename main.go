package main

import(
	"net/http"
	"fmt"
	"github.com/astaxie/beego/logs"
)

// Items func to get a file tail 5 row to return
func Items(w http.ResponseWriter, r *http.Request) {
	cmdFmt := `tail -5 %s`
	cmd := fmt.Sprintf(cmdFmt, appConf.listenFile)
	s, e := runShell(cmd)
	var ret string
	if len(s) != 0 {
		ret = s
	}else{
		ret = e
	}
	fmt.Println("content:", ret)
	fmt.Fprintf(w, ret)
}

func init() {
	file := "./conf/app.cfg"
	err := initConfig(file)
	if err != nil {
		return
	}

	err = initLogs(appConf.log)
	if err != nil {
		logs.Error("init logs failed:%v", err)
		return
	}
	logs.Info("init logs success")
}

func main() {
	http.HandleFunc("/items", Items)
	err := http.ListenAndServe("0.0.0.0:9090", nil)
	if err != nil {
		fmt.Println("http listen failed:", err)
		return
	}
}