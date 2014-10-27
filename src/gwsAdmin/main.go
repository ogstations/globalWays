package main

import (
	_"gwsAdmin/routers"
	"github.com/astaxie/beego"
	"os"
	"gwsAdmin/controllers"
	"github.com/beego/i18n"
	"runtime"
)

func main() {

	//多核运行
	np := runtime.NumCPU()
	if np >= 2 {
		runtime.GOMAXPROCS(np - 1)
	}

	//init
	initLog()
	controllers.InitApp()
	initTemplateFunc()

	beego.Run()
}

// 初始化log
func initLog() {

	// log setting
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(4)

	if beego.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}

}

// 初始化模版函数
func initTemplateFunc() {
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)


}
