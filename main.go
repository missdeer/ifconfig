package main

import (
	"os"

	"github.com/beego/beego/v2/server/web"
	_ "github.com/missdeer/ifconfig/routers"
)

const (
	HostVar = "HTTP_ADDR"
	PortVar = "HTTP_PORT"
)

func main() {
	port := web.AppConfig.DefaultString("httpport", "8080")
	if portEnv := os.Getenv(PortVar); portEnv != "" {
		port = portEnv
	}

	host := web.AppConfig.DefaultString("httpaddr", "127.0.0.1")
	if hostEnv := os.Getenv(HostVar); hostEnv != "" {
		host = hostEnv
	}
	web.Run(host + ":" + port)
}
