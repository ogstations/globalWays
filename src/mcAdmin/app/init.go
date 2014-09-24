package app

import (
	"github.com/revel/revel"
	"path/filepath"
	"github.com/Unknwon/goconfig"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	revel.OnAppStart(loadApiUrl)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var (
	// 分发渠道api url
	ChannelTypeUrl string
	// 渠道ID api url
	ChannelIdUrl string
	// 新建 memberCard api url
	NewMemberCardUrl string
)

// api url
func loadApiUrl() {

	config_file, err := goconfig.LoadConfigFile(filepath.Join(revel.BasePath, "conf", "api.conf"))
	if err != nil {
		revel.ERROR.Fatalf("laod configure file return error: %v, exit.", err)
	}

	ChannelTypeUrl, _ = config_file.GetValue("api", "api.channelType.url")
	ChannelIdUrl, _ = config_file.GetValue("api", "api.channelId.url")
	NewMemberCardUrl, _ = config_file.GetValue("api", "api.newCard.url")
}
