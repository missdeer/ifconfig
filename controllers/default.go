package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/beego/beego/v2/server/web"
	"github.com/zu1k/nali/pkg/geoip"
	"github.com/zu1k/nali/pkg/ip2region"
	"github.com/zu1k/nali/pkg/ipip"
	"github.com/zu1k/nali/pkg/qqwry"
)

const (
	QQWryPath        = "qqwry.dat"
	ZXIPv6WryPath    = "zxipv6wry.db"
	GeoLite2CityPath = "GeoLite2-City.mmdb"
	IPIPFreePath     = "ipipfree.ipdb"
	Ip2RegionPath    = "ip2region.db"
)

var (
	geoip2Instance    *geoip.GeoIP
	qqwryInstance     *qqwry.QQwry
	ipipInstance      *ipip.IPIPFree
	ip2regionInstance *ip2region.Ip2Region
)

func init() {
	geoip2Instance, _ = geoip.NewGeoIP(GeoLite2CityPath)
	qqwryInstance, _ = qqwry.NewQQwry(QQWryPath)
	ipipInstance, _ = ipip.NewIPIP(IPIPFreePath)
	ip2regionInstance, _ = ip2region.NewIp2Region(Ip2RegionPath)
}

type MainController struct {
	web.Controller
}

func (this *MainController) QueryGeoip2(ip string) (string, error) {
	if geoip2Instance == nil {
		return "", errors.New("Geoip2 service not available")
	}
	res, err := geoip2Instance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (this *MainController) GetLocationFromGeoIP2() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	res, err := this.QueryGeoip2(ip)
	if err != nil {
		this.Data["Value"] = err.Error()
	} else {
		this.Data["Value"] = res
	}
	this.TplName = "value.tpl"
}

func (this *MainController) QueryQQWry(ip string) (string, error) {
	if qqwryInstance == nil {
		return "", errors.New("QQWry service not available")
	}
	res, err := qqwryInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (this *MainController) GetLocationFromQQWry() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	res, err := this.QueryQQWry(ip)
	if err != nil {
		this.Data["Value"] = err.Error()
	} else {
		this.Data["Value"] = res
	}
	this.TplName = "value.tpl"
}

func (this *MainController) QueryIPIPFree(ip string) (string, error) {
	if ipipInstance == nil {
		return "", errors.New("IPIP free service not available")
	}
	res, err := ipipInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (this *MainController) GetLocationFromIPIP() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	res, err := this.QueryIPIPFree(ip)
	if err != nil {
		this.Data["Value"] = err.Error()
	} else {
		this.Data["Value"] = res
	}
	this.TplName = "value.tpl"
}

func (this *MainController) QueryIP2Region(ip string) (string, error) {
	if ip2regionInstance == nil {
		return "", errors.New("IP2Region service not available")
	}
	res, err := ip2regionInstance.Find(ip)
	if err != nil {
		return "", err
	}
	return res.String(), nil
}

func (this *MainController) GetLocationFromIP2Region() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	res, err := this.QueryIP2Region(ip)
	if err != nil {
		this.Data["Value"] = err.Error()
	} else {
		this.Data["Value"] = res
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetForwarded() {
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetHost() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		this.Data["Value"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		this.Data["Value"] = value
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetIP() {
	ip := this.GetString("ip", this.Ctx.Input.IP())
	this.Data["Value"] = ip
	this.TplName = "value.tpl"
}

func (this *MainController) GetPort() {
	remote_addr := []byte(this.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	this.Data["Value"] = string(remote_addr[pos+1:])
	this.TplName = "value.tpl"
}

func (this *MainController) GetVia() {
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Via"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetMime() {
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetLang() {
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetCharset() {
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Charset"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetEncoding() {
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetUserAgent() {
	this.Data["Value"] = this.Ctx.Request.UserAgent()
	this.TplName = "value.tpl"
}

func (this *MainController) GetConnection() {
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["Connection"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetKeepAlive() {
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		this.Data["Value"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	this.TplName = "value.tpl"
}

func (this *MainController) GetAll() {
	this.Data["Email"] = web.AppConfig.DefaultString("email", "")
	this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
	ip := this.GetString("ip", this.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		this.Data["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		this.Data["Host"] = value
	}
	this.Data["IP"] = ip
	this.Data["Geoip2"], _ = this.QueryGeoip2(ip)
	this.Data["IPIP"], _ = this.QueryIPIPFree(ip)
	this.Data["QQWry"], _ = this.QueryQQWry(ip)
	this.Data["IP2Region"], _ = this.QueryIP2Region(ip)
	remote_addr := []byte(this.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	this.Data["Port"] = string(remote_addr[pos+1:])
	this.Data["Method"] = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		this.Data["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		this.Data["Mime"] = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		this.Data["Connection"] = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		this.Data["Via"] = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		this.Data["Charset"] = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		this.Data["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		this.Data["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		this.Data["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	this.Data["Referer"] = this.Ctx.Input.Refer()

	this.TplName = "all.tpl"
}

type ifconfig struct {
	Email      string
	QQWry      string
	Geoip2     string
	IPIP       string
	IP2Region  string
	UserAgent  string
	Host       string
	IP         string
	Port       string
	Method     string
	Encoding   string
	Mime       string
	Connection string
	Via        string
	Charset    string
	Keepalive  string
	Forwarded  string
	Lang       string
	Referer    string
}

func (this *MainController) GetAllXML() {
	thisData := ifconfig{}
	thisData.Email = web.AppConfig.DefaultString("email", "")
	thisData.UserAgent = this.Ctx.Request.UserAgent()

	ip := this.GetString("ip", this.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		thisData.Host = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		thisData.Host = value
	}

	thisData.IP = ip
	thisData.Geoip2, _ = this.QueryGeoip2(ip)
	thisData.IPIP, _ = this.QueryIPIPFree(ip)
	thisData.QQWry, _ = this.QueryQQWry(ip)
	thisData.IP2Region, _ = this.QueryIP2Region(ip)
	log.Println(thisData.IP2Region)
	remote_addr := []byte(this.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	thisData.Port = string(remote_addr[pos+1:])
	thisData.Method = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		thisData.Encoding = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		thisData.Mime = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		thisData.Connection = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		thisData.Via = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		thisData.Charset = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		thisData.Keepalive = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		thisData.Forwarded = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		thisData.Lang = this.Ctx.Request.Header["Accept-Language"][0]
	}
	thisData.Referer = this.Ctx.Input.Refer()

	this.Data["xml"] = thisData
	this.ServeXML()
}

func (this *MainController) GetAllJSON() {
	thisData := make(map[string]interface{})
	thisData["Email"] = web.AppConfig.DefaultString("email", "")
	thisData["UserAgent"] = this.Ctx.Request.UserAgent()
	ip := this.GetString("ip", this.Ctx.Input.IP())
	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		thisData["Host"] = ""
	} else {
		var value string
		for _, v := range names {
			value += fmt.Sprintf("%s\n", v)
		}
		thisData["Host"] = value
	}

	thisData["IP"] = ip
	thisData["Geoip2"], _ = this.QueryGeoip2(ip)
	thisData["IPIP"], _ = this.QueryIPIPFree(ip)
	thisData["QQWry"], _ = this.QueryQQWry(ip)
	thisData["IP2Region"], _ = this.QueryIP2Region(ip)
	remote_addr := []byte(this.Ctx.Request.RemoteAddr)
	pos := bytes.IndexByte(remote_addr, ':')
	thisData["Port"] = string(remote_addr[pos+1:])
	thisData["Method"] = this.Ctx.Request.Method
	if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
		thisData["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
	}
	if len(this.Ctx.Request.Header["Accept"]) > 0 {
		thisData["Mime"] = this.Ctx.Request.Header["Accept"][0]
	}
	if len(this.Ctx.Request.Header["Connection"]) > 0 {
		thisData["Connection"] = this.Ctx.Request.Header["Connection"][0]
	}
	if len(this.Ctx.Request.Header["Via"]) > 0 {
		thisData["Via"] = this.Ctx.Request.Header["Via"][0]
	}
	if len(this.Ctx.Request.Header["Charset"]) > 0 {
		thisData["Charset"] = this.Ctx.Request.Header["Charset"][0]
	}
	if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
		thisData["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
	}
	if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
		thisData["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
	}
	if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
		thisData["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
	}
	thisData["Referer"] = this.Ctx.Input.Refer()

	this.Data["json"] = thisData
	this.ServeJSON()
}

func (this *MainController) GetGeo() {
	if noweb := os.Getenv("NOWEB"); noweb == "1" {
		this.Abort("404")
		return
	}

	ip := this.GetString("ip", this.Ctx.Input.IP())

	this.Data["IP"] = ip
	this.Data["Geoip2"], _ = this.QueryGeoip2(ip)
	this.Data["IPIP"], _ = this.QueryIPIPFree(ip)
	this.Data["QQWry"], _ = this.QueryQQWry(ip)
	this.Data["IP2Region"], _ = this.QueryIP2Region(ip)

	if strings.Contains(this.Ctx.Request.UserAgent(), "curl") {
		this.TplName = "geo.tpl"
	} else {
		this.Data["BaseUrl"] = web.AppConfig.DefaultString("baseurl", "ifconfig.ismisv.com")
		this.Data["Email"] = web.AppConfig.DefaultString("email", "")
		this.Data["UserAgent"] = this.Ctx.Request.UserAgent()
		names, err := net.LookupAddr(ip)
		if err != nil || len(names) == 0 {
			this.Data["Host"] = ""
		} else {
			var value string
			for _, v := range names {
				value += fmt.Sprintf("%s\n", v)
			}
			this.Data["Host"] = value
		}
		remote_addr := []byte(this.Ctx.Request.RemoteAddr)
		pos := bytes.IndexByte(remote_addr, ':')
		this.Data["Port"] = string(remote_addr[pos+1:])
		this.Data["Method"] = this.Ctx.Request.Method
		if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
			this.Data["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
		}
		if len(this.Ctx.Request.Header["Accept"]) > 0 {
			this.Data["Mime"] = this.Ctx.Request.Header["Accept"][0]
		}
		if len(this.Ctx.Request.Header["Connection"]) > 0 {
			this.Data["Connection"] = this.Ctx.Request.Header["Connection"][0]
		}
		if len(this.Ctx.Request.Header["Via"]) > 0 {
			this.Data["Via"] = this.Ctx.Request.Header["Via"][0]
		}
		if len(this.Ctx.Request.Header["Charset"]) > 0 {
			this.Data["Charset"] = this.Ctx.Request.Header["Charset"][0]
		}
		if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
			this.Data["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
		}
		if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
			this.Data["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
		}
		if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
			this.Data["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
		}
		this.Data["Referer"] = this.Ctx.Input.Refer()

		this.TplName = "index.tpl"
	}
}

func (this *MainController) Get() {
	if noweb := os.Getenv("NOWEB"); noweb == "1" {
		this.Abort("404")
		return
	}
	ip := this.GetString("ip", this.Ctx.Input.IP())
	this.Data["IP"] = ip
	if strings.Contains(this.Ctx.Request.UserAgent(), "curl") {
		this.TplName = "iponly.tpl"
	} else {
		this.Data["BaseUrl"] = web.AppConfig.DefaultString("baseurl", "ifconfig.ismisv.com")
		this.Data["Email"] = web.AppConfig.DefaultString("email", "")
		this.Data["UserAgent"] = this.Ctx.Request.UserAgent()

		names, err := net.LookupAddr(ip)
		if err != nil || len(names) == 0 {
			this.Data["Host"] = ""
		} else {
			var value string
			for _, v := range names {
				value += fmt.Sprintf("%s\n", v)
			}
			this.Data["Host"] = value
		}
		this.Data["Geoip2"], _ = this.QueryGeoip2(ip)
		this.Data["IPIP"], _ = this.QueryIPIPFree(ip)
		this.Data["QQWry"], _ = this.QueryQQWry(ip)
		this.Data["IP2Region"], _ = this.QueryIP2Region(ip)
		remote_addr := []byte(this.Ctx.Request.RemoteAddr)
		pos := bytes.IndexByte(remote_addr, ':')
		this.Data["Port"] = string(remote_addr[pos+1:])
		this.Data["Method"] = this.Ctx.Request.Method
		if len(this.Ctx.Request.Header["Accept-Encoding"]) > 0 {
			this.Data["Encoding"] = this.Ctx.Request.Header["Accept-Encoding"][0]
		}
		if len(this.Ctx.Request.Header["Accept"]) > 0 {
			this.Data["Mime"] = this.Ctx.Request.Header["Accept"][0]
		}
		if len(this.Ctx.Request.Header["Connection"]) > 0 {
			this.Data["Connection"] = this.Ctx.Request.Header["Connection"][0]
		}
		if len(this.Ctx.Request.Header["Via"]) > 0 {
			this.Data["Via"] = this.Ctx.Request.Header["Via"][0]
		}
		if len(this.Ctx.Request.Header["Charset"]) > 0 {
			this.Data["Charset"] = this.Ctx.Request.Header["Charset"][0]
		}
		if len(this.Ctx.Request.Header["KeepAlive"]) > 0 {
			this.Data["Keepalive"] = this.Ctx.Request.Header["KeepAlive"][0]
		}
		if len(this.Ctx.Request.Header["X-Forwarded-For"]) > 0 {
			this.Data["Forwarded"] = this.Ctx.Request.Header["X-Forwarded-For"][0]
		}
		if len(this.Ctx.Request.Header["Accept-Language"]) > 0 {
			this.Data["Lang"] = this.Ctx.Request.Header["Accept-Language"][0]
		}
		this.Data["Referer"] = this.Ctx.Input.Refer()

		this.TplName = "index.tpl"
	}
}
