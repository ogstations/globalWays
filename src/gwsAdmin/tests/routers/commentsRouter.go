package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Index",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"App",
			`/app`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Main",
			`/main`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:IndexController"],
		beego.ControllerComments{
			"Left",
			`/left`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:PublicController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:PublicController"],
		beego.ControllerComments{
			"Message",
			`/message`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"LoginGet",
			`/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"LoginPost",
			`/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"Logout",
			`/logout`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"ListAdmin",
			`/admin`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"AddAdmin",
			`/admin`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"UpdateAdmin",
			`/admin/:adminId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:AdminController"],
		beego.ControllerComments{
			"DelAdmin",
			`/admin/:adminId`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"ListChannelType",
			`/channelType`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"GetChannelTypeById",
			`/channelType/:channelId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"UpdateChannelTypeById",
			`/channelType/:channelId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["gwsAdmin/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"NewChannelType",
			`/channelType`,
			[]string{"post"},
			nil})

}
