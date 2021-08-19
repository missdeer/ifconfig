package routers

import (
	"github.com/beego/beego"
	"github.com/missdeer/ifconfig/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/ip", &controllers.MainController{}, "get:GetIP")
	beego.Router("/host", &controllers.MainController{}, "get:GetHost")
	beego.Router("/ua", &controllers.MainController{}, "get:GetUserAgent")
	beego.Router("/port", &controllers.MainController{}, "get:GetPort")
	beego.Router("/keepalive", &controllers.MainController{}, "get:GetKeepAlive")
	beego.Router("/lang", &controllers.MainController{}, "get:GetLang")
	beego.Router("/connection", &controllers.MainController{}, "get:GetConnection")
	beego.Router("/encoding", &controllers.MainController{}, "get:GetEncoding")
	beego.Router("/via", &controllers.MainController{}, "get:GetVia")
	beego.Router("/mime", &controllers.MainController{}, "get:GetMime")
	beego.Router("/charset", &controllers.MainController{}, "get:GetCharset")
	beego.Router("/forwarded", &controllers.MainController{}, "get:GetForwarded")
	beego.Router("/all", &controllers.MainController{}, "get:GetAll")
	beego.Router("/all.xml", &controllers.MainController{}, "get:GetAllXML")
	beego.Router("/all.json", &controllers.MainController{}, "get:GetAllJSON")
}
