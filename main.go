package main

import (
	"github.com/astaxie/beego"
	_ "github.com/missdeer/ifconfig/routers"
	"os"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"
)

func main() {
	var port string
	if port = os.Getenv(PortVar); port == "" {
		port = "8080"
	}
	beego.Run(beego.HttpAddr + ":" + port)
}
