package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var Log *logs.BeeLogger = nil

func init() {
	Log = logs.NewLogger(10000)
	if beego.AppConfig.String("runmode") == "dev" {
		Log.SetLogger("console", "")
	}
	appname := beego.AppConfig.String("appname")
	Log.SetLogger("file", `{"filename":"logs/`+appname+`.log","maxlines":0,"maxsize":0,"daily":true,"maxdays":7}`)

	Log.Async(2000)
	Log.SetLogFuncCallDepth(3)
	Log.EnableFuncCallDepth(true)
}

//错误日志
func LogError(format string, v ...interface{}) {
	Log.EnableFuncCallDepth(true)
	Log.Error(format, v...)
}

func LogInfo(format string, v ...interface{}) {
	Log.EnableFuncCallDepth(false)
	Log.Info(format, v...)
}
