package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"],
		beego.ControllerComments{
			"Get",
			`/:cardId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:MemberCardController"],
		beego.ControllerComments{
			"GetQrCode",
			`/:cardId/qrcode`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Get",
			`/:channelId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"] = append(beego.GlobalControllerRouter["mcAPI/controllers:ChannelTypeController"],
		beego.ControllerComments{
			"Put",
			`/:channelId`,
			[]string{"put"},
			nil})

}
