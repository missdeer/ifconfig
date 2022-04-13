package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/missdeer/ifconfig/controllers"
)

func init() {
	web.Router("/", &controllers.MainController{}, "get:Get")
	web.Router("/ipip", &controllers.MainController{}, "get:GetLocationFromIPIP")
	web.Router("/ip2region", &controllers.MainController{}, "get:GetLocationFromIP2Region")
	web.Router("/qqwry", &controllers.MainController{}, "get:GetLocationFromQQWry")
	web.Router("/geoip2", &controllers.MainController{}, "get:GetLocationFromGeoIP2")
	web.Router("/ip", &controllers.MainController{}, "get:GetIP")
	web.Router("/host", &controllers.MainController{}, "get:GetHost")
	web.Router("/ua", &controllers.MainController{}, "get:GetUserAgent")
	web.Router("/port", &controllers.MainController{}, "get:GetPort")
	web.Router("/keepalive", &controllers.MainController{}, "get:GetKeepAlive")
	web.Router("/lang", &controllers.MainController{}, "get:GetLang")
	web.Router("/connection", &controllers.MainController{}, "get:GetConnection")
	web.Router("/encoding", &controllers.MainController{}, "get:GetEncoding")
	web.Router("/via", &controllers.MainController{}, "get:GetVia")
	web.Router("/mime", &controllers.MainController{}, "get:GetMime")
	web.Router("/charset", &controllers.MainController{}, "get:GetCharset")
	web.Router("/forwarded", &controllers.MainController{}, "get:GetForwarded")
	web.Router("/all", &controllers.MainController{}, "get:GetAll")
	web.Router("/all.xml", &controllers.MainController{}, "get:GetAllXML")
	web.Router("/all.json", &controllers.MainController{}, "get:GetAllJSON")
}
