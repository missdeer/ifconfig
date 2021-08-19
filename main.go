package main

import (
	"os"

	"github.com/beego/beego"
	_ "github.com/missdeer/ifconfig/routers"
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
	beego.Run(beego.BConfig.Listen.HTTPAddr + ":" + port)
}
