package main

import (
	"os"

	"github.com/fatedier/frp/pkg/utils/log"
	"github.com/fatedier/frp/pkg/utils/conn"
)

func main() {
	err := LoadConf("./frps.ini")
	if err != nil {
		os.Exit(-1)
	}

	log.InitLog(LogWay, LogFile, LogLevel)

	l, err := conn.Listen(BindAddr, BindPort)
	if err != nil {
		log.Error("Create listener error, %v", err)
		os.Exit(-1)
	}

	log.Info("Start frps success")
	ProcessControlConn(l)
}
