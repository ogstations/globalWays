package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/memberCard","description":"会员卡API\n"},{"path":"/channelType","description":"分发渠道API\n"}],"info":{"title":"环途国际会员卡系统 API","description":"会员卡API，用于GlobalWays会员卡系统构建","contact":"mint.zhao.chiu@gmail.com","termsOfServiceUrl":"http://www.globalways.com/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
var subapi string = `{"/channelType":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/channelType","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"createChannelType","type":"","summary":"创建一个分发渠道","parameters":[{"paramType":"body","name":"channelType","description":"\"分发渠道json\"","dataType":"ChannelType","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.ChannelType","responseModel":"ChannelType"},{"code":403,"message":"body is empty","responseModel":""},{"code":500,"message":"json unmarshal error","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"GetChannels","type":"","summary":"获取所有的分发渠道","responseMessages":[{"code":200,"message":"models.ChannelType","responseModel":"ChannelType"}]}]},{"path":"/:channelId","description":"","operations":[{"httpMethod":"GET","nickname":"GetChannel","type":"","summary":"通过ID获取分发渠道","parameters":[{"paramType":"path","name":"channelId","description":"\"分发渠道ID\"","dataType":"int64","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.ChannelType","responseModel":"ChannelType"},{"code":403,"message":":channelId is empty","responseModel":""}]}]},{"path":"/:channelId","description":"","operations":[{"httpMethod":"PUT","nickname":"updateChannel","type":"","summary":"更新分发渠道信息","parameters":[{"paramType":"path","name":"channelId","description":"\"想更新的渠道ID\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"channelType","description":"\"更新的渠道信息\"","dataType":"ChannelType","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.ChannelType","responseModel":"ChannelType"},{"code":403,"message":":channelId is not int","responseModel":""}]}]}],"models":{"ChannelType":{"id":"ChannelType","properties":{"ChannelDesc":{"type":"string","description":"","format":""},"ChannelName":{"type":"string","description":"","format":""},"Id":{"type":"int64","description":"","format":""}}}}},"/memberCard":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/memberCard","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"createMemberCards","type":"","summary":"批量生成会员卡","parameters":[{"paramType":"body","name":"reqMemberCards","description":"\"请求新建会员卡参数\"","dataType":"ReqNewMemberCards","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MemberCard.Id","responseModel":""},{"code":400,"message":"bad request","responseModel":""},{"code":403,"message":"body is empty","responseModel":""},{"code":500,"message":"internal server error","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"getMemberCards","type":"","summary":"获取会员卡号列表","parameters":[{"paramType":"query","name":"page","description":"\"分页page\"","dataType":"int64","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0},{"paramType":"query","name":"size","description":"\"分页size\"","dataType":"int64","type":"","format":"","allowMultiple":false,"required":false,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MemberCard","responseModel":""},{"code":500,"message":"internal server error","responseModel":""}]}]},{"path":"/:cardId","description":"","operations":[{"httpMethod":"GET","nickname":"getMemberCard","type":"","summary":"通过ID获取会员卡信息","parameters":[{"paramType":"path","name":"cardId","description":"\"会员卡ID\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MemberCard","responseModel":""},{"code":400,"message":"bad request","responseModel":""},{"code":500,"message":"internal server error","responseModel":""}]}]},{"path":"/:cardId/qrcode","description":"","operations":[{"httpMethod":"GET","nickname":"getMemberCardQrCode","type":"","summary":"通过ID获取会员卡二维码","parameters":[{"paramType":"path","name":"cardId","description":"\"会员卡ID\"","dataType":"int","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"qrCode","responseModel":""},{"code":400,"message":"bad request","responseModel":""},{"code":500,"message":"internal server error","responseModel":""}]}]}]}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
