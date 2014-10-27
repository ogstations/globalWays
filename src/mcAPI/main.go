package main

import (
	_ "mcAPI/docs"
	_ "mcAPI/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"os"
)

func main() {
	beego.DirectoryIndex = true
	beego.StaticDir["/swagger"] = "swagger"

	beego.Run()
}

func init() {

	corsHandler := func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
		ctx.Output.Header("Access-Control-Allow-Methods", "*")
	}

	beego.InsertFilter("*", beego.BeforeRouter, corsHandler)

	// log setting
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(4)
	if beego.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}
}
